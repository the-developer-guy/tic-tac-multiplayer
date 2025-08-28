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
	Token          string
	Name           string
	isBanned       bool
	dateofRegister string
}

type Lobby struct {
	Players map[string]string `json:"players"`
	LobbyID string            `json:"lobbyID"`
	Grid    *TicTacToeGrid    `json:"gameGrid"`
	lock    sync.Mutex        // TODO add access methods to Lobby
}

func NewLobby(token1 string, token2 string) *Lobby {
	return &Lobby{
		Players: map[string]string{
			MarkX.String(): token1,
			MarkO.String(): token2,
		},
		LobbyID: uuid.NewString(),
		Grid:    NewTicTacToeGrid(),
	}
}
