package server

import (
	"encoding/json"
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

	dataStore = NewFetchedData() //Temporary, it has to be changed as soon DB is implemented.

)

type LoginStruct struct {
	LoginFailed bool
}

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

	data := LoginStruct{LoginFailed: loginFailed}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CheckSession(w http.ResponseWriter, req *http.Request) error {
	session, _ := store.Get(req, sessionName)
	auth, ok := session.Values["authenticated"].(bool)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return fmt.Errorf("Forbidden")
	}
	return nil
}

func (ttts *TicTacToeServer) accessControl(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	username := req.Form.Get("user")
	password := req.Form.Get("password")

	adminUser := envFile["ADMIN_USER"]
	adminPass := envFile["ADMIN_PASS"]
	if adminUser == "" || adminPass == "" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	okUser := adminUser == username
	okPass := adminPass == password

	session, _ := store.Get(req, sessionName)

	if !okUser || !okPass {
		session.AddFlash("invalid_credentials", "login")
		_ = session.Save(req, w)

		http.Redirect(w, req, "/login/", http.StatusSeeOther)
		return
	}
	session.Values["authenticated"] = true
	if err := session.Save(req, w); err != nil {
		fmt.Printf("session save error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, "/adminpage/", http.StatusSeeOther)
	return
}

func (ttts *TicTacToeServer) adminPage(w http.ResponseWriter, req *http.Request) {
	if err := CheckSession(w, req); err != nil {
		return
	}

	t, _ := template.ParseFiles("./templates/interface.html")
	t.Execute(w, nil)
}

func (ttts *TicTacToeServer) GetData(w http.ResponseWriter, req *http.Request) {
	if err := CheckSession(w, req); err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	allPlayers := dataStore.GetAllData()

	json.NewEncoder(w).Encode(allPlayers)
}

func (ttts *TicTacToeServer) NewPlayer(w http.ResponseWriter, req *http.Request) {
	if err := CheckSession(w, req); err != nil {
		return
	}

	req.ParseForm()
	name := req.Form.Get("name")
	if name == "" {
		http.Error(w, "Missing playerName parameter", http.StatusBadRequest)
		return
	}

	dataStore.NewPlayer(name)
	http.Redirect(w, req, "/adminpage/", http.StatusSeeOther)
}

func (ttts *TicTacToeServer) RemoveUser(w http.ResponseWriter, req *http.Request) {
	if err := CheckSession(w, req); err != nil {
		return
	}

	req.ParseForm()
	token := req.Form.Get("token")
	if token == "" {
		http.Error(w, "Missing Token parameter", http.StatusBadRequest)
		return
	}
	dataStore.RemovePlayer(token)
	w.WriteHeader(http.StatusOK)
}

func (ttts *TicTacToeServer) RegenerateToken(w http.ResponseWriter, req *http.Request) {
	if err := CheckSession(w, req); err != nil {
		return
	}

	req.ParseForm()
	token := req.Form.Get("token")
	if token == "" {
		http.Error(w, "Missing Token parameter", http.StatusBadRequest)
		return
	}

	dataStore.RegenerateToken(token)
	w.WriteHeader(http.StatusOK)
}

func (ttts *TicTacToeServer) EditPlayerPermissions(w http.ResponseWriter, req *http.Request) {
	if err := CheckSession(w, req); err != nil {
		return
	}

	req.ParseForm()
	token := req.Form.Get("token")
	if token == "" {
		http.Error(w, "Missing Token parameter", http.StatusBadRequest)
		return
	}
	dataStore.HandlePlayerAccess(token)
	w.WriteHeader(http.StatusOK)

}
