package server

import (
	"sync"
	"time"

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

type PlayerScores struct {
	WinCount  int `json:"winCount"`
	LoseCount int `json:"loseCount"`
	TieCount  int `json:"tieCount"`
}

type Player struct {
	Name           string        `json:"name"`
	Token          string        `json:"token"`
	IsBanned       *time.Time    `json:"isBanned"`
	DateOfRegister string        `json:"dateOfRegister"`
	Scores         *PlayerScores `json:"scores"`
}

type Lobby struct {
	PlayerAMark string `json:"playerAMark"`
	PlayerBMark string `json:"playerBMark"`

	PlayerAToken string `json:"playerAToken"`
	PlayerBToken string `json:"playerBToken"`

	LobbyID string         `json:"lobbyID"`
	Grid    *TicTacToeGrid `json:"gameGrid"`
	lock    sync.Mutex     // TODO add access methods to Lobby
}

func NewLobby(token1 string, token2 string) *Lobby {
	return &Lobby{
		PlayerAMark: MarkX.String(),
		PlayerBMark: MarkO.String(),

		PlayerAToken: token1,
		PlayerBToken: token2,

		LobbyID: uuid.NewString(),
		Grid:    NewTicTacToeGrid(),
	}
}
