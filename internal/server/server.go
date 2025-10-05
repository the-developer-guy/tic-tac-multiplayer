package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"sync"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal"
	"github.com/the-developer-guy/tic-tac-multiplayer/internal/auth"
	"github.com/the-developer-guy/tic-tac-multiplayer/internal/game"
)

type GameServer struct {
	settings                *internal.AppConfig
	TournamentLobbies       map[string]*game.Lobby
	ActiveTournamentLobbies map[string]*game.Lobby
	ScheduledLobbies        map[string]*game.Lobby
	ReadyPlayerIDs          []int64
	lobbiesLock             sync.Mutex
	playersLock             sync.Mutex
	auth                    *auth.UserAuth
	players                 *auth.PlayerAuth
}

func NewGameServer(ac *internal.AppConfig) *GameServer {
	gs := GameServer{
		settings:                ac,
		TournamentLobbies:       make(map[string]*game.Lobby),
		ActiveTournamentLobbies: make(map[string]*game.Lobby),
		ScheduledLobbies:        make(map[string]*game.Lobby),
		auth:                    auth.NewUserAuth(),
		players:                 auth.NewPlayerAuth(),
	}
	gs.auth.AddUser(ac.AdminUser, ac.AdminPassword)

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
	_, lobbyExists := gs.TournamentLobbies[lobby.LobbyID]
	if lobbyExists {
		return fmt.Errorf("lobby ID %s already exists", lobby.LobbyID)
	}

	_, lobbyActive := gs.ActiveTournamentLobbies[lobby.LobbyID]
	if lobbyActive {
		return fmt.Errorf("lobby ID %s already active", lobby.LobbyID)
	}

	gs.lobbiesLock.Lock()
	gs.TournamentLobbies[lobby.LobbyID] = lobby
	gs.ActiveTournamentLobbies[lobby.LobbyID] = lobby
	gs.lobbiesLock.Unlock()

	return nil
}

func (gs *GameServer) GetLobby(lobbyId string) (*game.Lobby, error) {
	l, ok := gs.TournamentLobbies[lobbyId]
	if !ok {
		return nil, fmt.Errorf("no lobby ID %s", lobbyId)
	}

	return l, nil
}

func (gs *GameServer) GetReadyLobby(playerId int64) (*game.Lobby, error) {
	gs.lobbiesLock.Lock()
	var lobby *game.Lobby
	for _, l := range gs.ActiveTournamentLobbies {
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

	playerIndex := slices.Index(gs.ReadyPlayerIDs, id)
	if playerIndex == -1 {
		gs.ReadyPlayerIDs = append(gs.ReadyPlayerIDs, id)
	}

	gs.playersLock.Unlock()

	if playerIndex != -1 {
		return fmt.Errorf("player ID %d already scheduled for a match", id)
	}

	return nil
}

func (gs *GameServer) Json() ([]byte, error) {
	gs.lobbiesLock.Lock()
	payload, err := json.Marshal(gs.TournamentLobbies)
	gs.lobbiesLock.Unlock()

	return payload, err
}
