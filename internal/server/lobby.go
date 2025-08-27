package server

import (
	"sync"

	"github.com/google/uuid"
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
		return "?"
	}
}

type Player struct {
	Token string `json:"token"`
	Mark  Mark   `json:"mark"`
}

type Lobby struct {
	Players *[]Player      `json:"players"`
	LobbyID string         `json:"lobbyID"`
	Grid    *TicTacToeGrid `json:"gameGrid"`
	lock    sync.Mutex     // TODO add access methods to Lobby
}

func NewLobby(token1 string, token2 string) *Lobby {
	return &Lobby{
		Players: &[]Player{
			{Token: token1, Mark: MarkX},
			{Token: token2, Mark: MarkO},
		},
		LobbyID: uuid.NewString(),
		Grid:    NewTicTacToeGrid(),
	}
}
