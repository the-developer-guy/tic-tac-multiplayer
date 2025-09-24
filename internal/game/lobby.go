package game

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type Mark int

const (
	MarkEmpty Mark = iota
	MarkX
	MarkO
)

func (m Mark) String() string {
	switch m {
	case MarkEmpty:
		return ""
	case MarkX:
		return "X"
	case MarkO:
		return "O"
	default:
		return "?"
	}
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

func (l *Lobby) PlaceMark(token string, x, y int) error {
	return errors.New("not implemented")
}
