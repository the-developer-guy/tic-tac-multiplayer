package server

import (
	"math/rand"
	"net/http"
	"time"
)

type Player struct {
	Token string
	Mark  string
}

type Lobby struct {
	Players []Player      `json:"players"`
	LobbyID string        `json:"lobbyid"`
	Grid    TicTacToeGrid `json:"gamegrid"`
}

func ValidatePOST(w http.ResponseWriter, req *http.Request, params []string) bool {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	req.ParseForm()

	for _, param := range params {
		if len(param) == 0 {
			return false
		}
	}
	if len(params[0]) == 0 || len(params[1]) == 0 {
		http.Error(w, "Missing Arguments", http.StatusBadRequest)
		return false
	}

	return true
}

func CreateLobby(req *http.Request) Lobby {
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
		Grid:    GenerateGrid(),
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
