package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/lpernett/godotenv"
)

var (
	authKey = securecookie.GenerateRandomKey(64) // HMAC
	encKey  = securecookie.GenerateRandomKey(32) // AES-256
	store   = sessions.NewCookieStore(authKey, encKey)

	sessionName = "ttt_session"
	envFile, _  = godotenv.Read(".env")
)

func (ttts *TicTacToeServer) LoginPage(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := store.Get(req, sessionName)
	flashes := session.Flashes("login")
	_ = session.Save(req, w) // It removes the flashes after reading.

	var loginFailed bool
	if len(flashes) > 0 {
		loginFailed = true
	}

	data := struct {
		LoginFailed bool
	}{
		LoginFailed: loginFailed,
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ttts *TicTacToeServer) accessControl(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	username := req.Form.Get("usr")
	password := req.Form.Get("pass")

	adminUser := envFile["ADMIN_USER"]
	adminPass := envFile["ADMIN_PASS"]
	if adminUser == "" || adminPass == "" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	okUser := adminUser == username
	okPass := adminPass == password

	session, _ := store.Get(req, sessionName)

	if okUser && okPass {
		session.Values["authenticated"] = true
		if err := session.Save(req, w); err != nil {
			fmt.Printf("session save error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/ainterface/", http.StatusSeeOther)
		return
	}

	session.AddFlash("invalid_credentials", "login")
	_ = session.Save(req, w)
	//fmt.Println(session.Flashes())
	http.Redirect(w, req, "/login/", http.StatusSeeOther)
}

func (ttts *TicTacToeServer) ainterFace(w http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, sessionName)
	auth, ok := session.Values["authenticated"].(bool)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	dataStore := NewFetchedData()
	allPlayers := dataStore.GetAllData()
	t, _ := template.ParseFiles("./templates/interface.html")
	t.Execute(w, allPlayers)
}
