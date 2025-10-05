package game

import (
	"fmt"
	"sync"
	"time"

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
	PlayerAMark string `json:"playerAMark"` // Mark O by default
	PlayerBMark string `json:"playerBMark"` // Mark X by default

	PlayerAToken string `json:"playerAToken"`
	PlayerBToken string `json:"playerBToken"`

	LobbyID string         `json:"lobbyID"`
	Grid    *TicTacToeGrid `json:"gameGrid"`
	lock    sync.Mutex     // TODO add access methods to Lobby

	StartTime time.Time
}

func NewLobby(token1 string, token2 string, scheduledStart time.Time) *Lobby {
	l := Lobby{
		PlayerAMark: MarkX.String(),
		PlayerBMark: MarkO.String(),

		PlayerAToken: token1,
		PlayerBToken: token2,

		LobbyID: uuid.NewString(),
		Grid:    NewTicTacToeGrid(),

		StartTime: scheduledStart,
	}

	return &l
}

func (l *Lobby) PlaceMark(x, y int, token string) error {
	switch token {
	case l.PlayerAToken:
		err := l.Grid.PlaceMark(x, y, MarkO)
		if err != nil {
			return err
		}
	case l.PlayerBToken:
		err := l.Grid.PlaceMark(x, y, MarkX)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("wrong token for lobby %s", l.LobbyID)
	}

	return nil
}

func (l *Lobby) ScheduleJson() []byte {
	jsonString := fmt.Sprintf("{\"lobbyId\": \"%s\", \"nextGame\": %d}", l.LobbyID, l.StartTime.UnixMilli())

	return []byte(jsonString)
}
