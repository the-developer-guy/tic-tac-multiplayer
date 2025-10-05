package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal"
	"github.com/the-developer-guy/tic-tac-multiplayer/internal/auth"
	"github.com/the-developer-guy/tic-tac-multiplayer/internal/game"
)

type GameServer struct {
	settings    *internal.AppConfig
	Lobbies     map[string]*game.Lobby
	lobbiesLock sync.Mutex
	auth        *auth.UserAuth
}

func NewGameServer(ac *internal.AppConfig) *GameServer {
	gs := GameServer{
		settings: ac,
		Lobbies:  make(map[string]*game.Lobby),
		auth:     auth.NewUserAuth(),
	}
	gs.auth.AddUser(ac.AdminUser, ac.AdminPassword)

	if ac.LocalTest {
		l := game.NewLobby("test", "server")
		gs.AddLobby(l)
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

func (gs *GameServer) AddLobby(lobby *game.Lobby) error {
	_, lobbyExists := gs.Lobbies[lobby.LobbyID]
	if lobbyExists {
		return fmt.Errorf("lobby ID %s already exists", lobby.LobbyID)
	}

	gs.lobbiesLock.Lock()
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

func (gs *GameServer) Json() ([]byte, error) {
	gs.lobbiesLock.Lock()
	payload, err := json.Marshal(gs.Lobbies)
	gs.lobbiesLock.Unlock()

	return payload, err
}
