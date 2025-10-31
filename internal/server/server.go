package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"math/rand"
	"net/http"
	"slices"
	"sync"
	"time"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal"
	"github.com/the-developer-guy/tic-tac-multiplayer/internal/auth"
	"github.com/the-developer-guy/tic-tac-multiplayer/internal/game"
)

type ScheduledPlayer struct {
	Player    *auth.Player
	LobbyID   string
	ReadyTime int64
}

type GameServer struct {
	settings         *internal.AppConfig
	Lobbies          map[string]*game.Lobby
	ScheduledLobbies map[string]*game.Lobby
	ReadyPlayers     map[int64]*ScheduledPlayer
	lobbiesLock      sync.Mutex
	playersLock      sync.Mutex
	auth             *auth.UserAuth
	players          *auth.PlayerAuth
}

func NewGameServer(ac *internal.AppConfig) *GameServer {
	gs := GameServer{
		settings:         ac,
		Lobbies:          make(map[string]*game.Lobby),
		ScheduledLobbies: make(map[string]*game.Lobby),
		ReadyPlayers:     make(map[int64]*ScheduledPlayer),
		auth:             auth.NewUserAuth(),
		players:          auth.NewPlayerAuth(),
	}
	gs.auth.AddUser(ac.AdminUser, ac.AdminPassword)

	if gs.settings.LocalTest {
		p := auth.NewPlayer("server", "")
		gs.players.AddPlayer(0, p)
	}

	return &gs
}

func (gs *GameServer) RegisterApiHandles() {
	http.HandleFunc("GET /playerinfo/{playerId}/", gs.HandlePlayerInfo)
	http.HandleFunc("GET /ready/{playerId}/", gs.HandleReadyPlayer)

	http.HandleFunc("GET /getgrid/{lobbyId}/", gs.HandleGetGameGrid)
	http.HandleFunc("POST /place/{lobbyId}/", gs.HandlePlaceMark)
	http.HandleFunc("GET /getscores/", gs.HandleFetchPlayerScores)
	http.HandleFunc("GET /scores/", gs.HandleScoresView)
}

func (gs *GameServer) RegisterAdminHandles() {
	http.HandleFunc("GET /admin/players/", gs.HandleAdminListPlayers)

	http.HandleFunc("/login/", gs.HandleLoginView)
	http.HandleFunc("POST /accessc", gs.HandleAccessControl)
	http.HandleFunc("/adminpage/", gs.HandleAdminView)
	http.HandleFunc("/fetchdata/", gs.HandleGetData)
	http.HandleFunc("POST /manualnewplayer/", gs.HandleNewPlayer)
	http.HandleFunc("POST /regeneratetoken/", gs.HandleRegenerateToken)
	http.HandleFunc("POST /handleplayeraccess/", gs.HandleEditPlayerPermissions)
}

func (gs *GameServer) ScheduleTournamentMatch(player1Token, player2Token string) error {
	return errors.New("not implemented")
}

func (gs *GameServer) AddLobby(lobby *game.Lobby) error {
	_, lobbyExists := gs.Lobbies[lobby.LobbyID]
	if lobbyExists {
		return fmt.Errorf("lobby ID %s already exists", lobby.LobbyID)
	}

	_, lobbyActive := gs.Lobbies[lobby.LobbyID]
	if lobbyActive {
		return fmt.Errorf("lobby ID %s already active", lobby.LobbyID)
	}

	gs.lobbiesLock.Lock()
	gs.Lobbies[lobby.LobbyID] = lobby
	gs.Lobbies[lobby.LobbyID] = lobby
	gs.lobbiesLock.Unlock()

	return nil
}

func (gs *GameServer) GetLobby(lobbyId string) (*game.Lobby, error) {
	l, ok := gs.Lobbies[lobbyId]
	if !ok {
		return nil, fmt.Errorf("no lobby ID %s", lobbyId)
	}

	return l, nil
}

func (gs *GameServer) GetReadyLobby(playerId int64) (*game.Lobby, error) {
	gs.lobbiesLock.Lock()
	var lobby *game.Lobby
	for _, l := range gs.Lobbies {
		if playerId == l.PlayerAId || playerId == l.PlayerBId {
			lobby = l
			break
		}
	}
	gs.lobbiesLock.Unlock()

	return lobby, nil
}

func (gs *GameServer) AddReadyPlayer(id int64) error {
	gs.playersLock.Lock()

	_, ok := gs.ReadyPlayers[id]

	if !ok {
		p, err := gs.players.GetPlayer(id)
		if err != nil {
			gs.playersLock.Unlock()
			return err
		}
		gs.ReadyPlayers[id] = &ScheduledPlayer{
			Player:    p,
			ReadyTime: time.Now().UnixMilli(),
		}

	}

	gs.playersLock.Unlock()

	if ok {
		return fmt.Errorf("player ID %d already scheduled for a match", id)
	}

	return nil
}

func (gs *GameServer) ScheduleGame() error {

	ids := slices.Collect(maps.Keys(gs.ReadyPlayers))
	if len(ids) < 2 {
		return errors.New("not enough ready players to schedule")
	} //else if len(ids) < 4 {
	//return errors.New("not enough ready players to properly schedule")
	//}

	playerAId := ids[rand.Intn(len(ids))]
	playerBId := ids[rand.Intn(len(ids))]
	for playerBId == playerAId {
		playerBId = ids[rand.Intn(len(ids))]
	}

	PlayerA, errA := gs.players.GetPlayer(playerAId)
	if errA != nil {
		return errA
	}
	PlayerB, errB := gs.players.GetPlayer(playerBId)
	if errB != nil {
		return errA
	}

	l := game.NewLobby(PlayerA.Token, PlayerB.Token, playerAId, playerBId, time.Now())
	gs.AddLobby(l)

	delete(gs.ReadyPlayers, playerAId)
	delete(gs.ReadyPlayers, playerBId)

	return nil
}

func (gs *GameServer) Json() ([]byte, error) {
	gs.lobbiesLock.Lock()
	payload, err := json.Marshal(gs.Lobbies)
	gs.lobbiesLock.Unlock()

	return payload, err
}
