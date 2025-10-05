package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/the-developer-guy/tic-tac-multiplayer/internal/auth"
)

var (
	authKey = securecookie.GenerateRandomKey(64) // HMAC
	encKey  = securecookie.GenerateRandomKey(32) // AES-256
	store   = sessions.NewCookieStore(authKey, encKey)

	sessionName = "ttt_session"
	dataStore   = NewFetchedData() //Temporary, it has to be changed as soon DB is implemented.

)

type FetchScores struct {
	Name  string            `json:"name"`
	Score auth.PlayerScores `json:"scores"`
}

type LoginStruct struct {
	LoginFailed bool
}

func (gs *GameServer) HandleLoginView(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, _ := store.Get(req, sessionName)
	flashes := session.Flashes("login")
	session.Save(req, w) // It removes the flashes after reading.

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
	session, err := store.Get(req, sessionName)
	if err != nil {
		fmt.Printf("failed to save session in accessControl: %v\n", err)
	}

	auth, ok := session.Values["authenticated"].(bool)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return fmt.Errorf("Forbidden")
	}
	return nil
}

func (gs *GameServer) HandleAccessControl(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	username := req.Form.Get("user")
	password := req.Form.Get("password")

	session, _ := store.Get(req, sessionName)

	user, err := gs.auth.GetUser(username)
	if err != nil {
		session.AddFlash("invalid_credentials", "login")
		_ = session.Save(req, w)

		http.Redirect(w, req, "/login/", http.StatusSeeOther)
		return
	}

	if !user.CheckPassword(password) {
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
}

func (gs *GameServer) HandleAdminView(w http.ResponseWriter, req *http.Request) {
	if err := CheckSession(w, req); err != nil {
		return
	}

	t, _ := template.ParseFiles("./templates/interface.html")
	t.Execute(w, nil)
}

func (gs *GameServer) HandleGetData(w http.ResponseWriter, req *http.Request) {
	if err := CheckSession(w, req); err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	allPlayers := dataStore.GetAllData()

	json.NewEncoder(w).Encode(allPlayers)
}

func (gs *GameServer) HandleNewPlayer(w http.ResponseWriter, req *http.Request) {
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

func (gs *GameServer) HandleRegenerateToken(w http.ResponseWriter, req *http.Request) {
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

func (gs *GameServer) HandleEditPlayerPermissions(w http.ResponseWriter, req *http.Request) {
	if err := CheckSession(w, req); err != nil {
		return
	}

	req.ParseForm()
	token := req.Form.Get("token")
	if token == "" {
		http.Error(w, "Missing Token parameter", http.StatusBadRequest)
		return
	}
	dataStore.ValidatePlayerAccess(token)
	w.WriteHeader(http.StatusOK)

}

func (gs *GameServer) HandleFetchPlayerScores(w http.ResponseWriter, req *http.Request) {
	r := dataStore.GetScores()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

func (gs *GameServer) HandleScoresView(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./templates/scores.html")
	t.Execute(w, nil)
}
