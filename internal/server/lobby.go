package server

import (
	"fmt"
	"net/http"
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
	LobbyID string        `json:"lobbyID"`
	Grid    TicTacToeGrid `json:"gameGrid"`
}

// Constructor
func NewLobby(token1 string, lenOfLobbies int) Lobby {
	return Lobby{
		Players: []Player{
			{Token: token1, Mark: MarkX},
			{Token: "", Mark: MarkO},
		},
		LobbyID: fmt.Sprintf("lobby_%d", lenOfLobbies),
		Grid:    GenerateGrid(),
	}
}

func ReturnLobby(req *http.Request, lenOfLobbies int) Lobby {
	req.ParseForm()
	token1 := req.Form.Get("token")

	return NewLobby(token1, lenOfLobbies)
}
