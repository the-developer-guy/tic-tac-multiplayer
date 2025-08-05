package server

import (
	"math/rand"
	"net/http"
	"time"
)

type Mark int

const (
	MarkX Mark = iota
	MarkO
)

func (m Mark) String() string {
	switch m {
	case MarkX:
		return "X"
	case MarkO:
		return "O"
	default:
		return ""
	}
}

type Player struct {
	Token string `json:"token"`
	Mark  Mark   `json:"mark"`
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

// Constructor
func NewLobby(token1, token2 string) Lobby {
	return Lobby{
		Players: []Player{
			{Token: token1, Mark: MarkX},
			{Token: token2, Mark: MarkO},
		},
		LobbyID: randomID(),
		Grid:    GenerateGrid(),
	}
}

func LobbyResponse(req *http.Request) Lobby {
	req.ParseForm()
	token1 := req.Form.Get("token1")
	token2 := req.Form.Get("token2")
	return NewLobby(token1, token2)
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
