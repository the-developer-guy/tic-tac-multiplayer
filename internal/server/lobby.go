package server

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Player struct {
	Token string //`json: "token" `
	Mark  string //`json: "mark" `
}

type Lobby struct {
	Players []Player      `json:"players"`
	LobbyID string        `json:"lobbyid"`
	Grid    TicTacToeGrid `json:"gamegrid"`
}

func CreateLobby(w http.ResponseWriter, req *http.Request) bool {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	req.ParseForm()
	params := []string{
		req.Form.Get("token1"),
		req.Form.Get("token2"),
	}

	if len(params[0]) == 0 || len(params[1]) == 0 {
		http.Error(w, "Missing Arguments", http.StatusBadRequest)
		return false
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Success")

	return true
}

func PrepareLobby(req *http.Request) Lobby {
	req.ParseForm()
	params := []string{
		req.Form.Get("token1"),
		req.Form.Get("token2"),
	}

	lobby := Lobby{
		Players: []Player{
			{Token: params[0], Mark: "X"},
			{Token: params[1], Mark: "O"},
		},
		LobbyID: randomID(),
		Grid:    writeGridJson(),
	}
	return lobby
}

func randomID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, 7)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}

	return string(result)
}
