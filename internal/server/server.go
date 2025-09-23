package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal"
	"github.com/the-developer-guy/tic-tac-multiplayer/internal/auth"
)

type TicTacToeServer struct {
	settings    *internal.AppConfig
	Lobbies     map[string]*Lobby
	lobbiesLock sync.Mutex
	auth        auth.UserAuth
}

func NewTicTacToeServer(ac *internal.AppConfig) *TicTacToeServer {
	ttts := TicTacToeServer{
		settings: ac,
		Lobbies:  make(map[string]*Lobby),
	}
	ttts.auth.AddUser(ac.AdminUser, ac.AdminPassword)

	return &ttts
}

func (ttts *TicTacToeServer) RegisterApiHandles() {
	http.HandleFunc("GET /playerinfo/{playerId}/", ttts.HandlePlayerInfo)
	http.HandleFunc("GET /ready/{playerId}/", ttts.HandleReadyPlayer)

	http.HandleFunc("GET /getgrid/{lobbyId}/", ttts.HandleGetGameGrid)
	http.HandleFunc("POST /place/{lobbyId}/", ttts.HandlePlaceMark)
	http.HandleFunc("GET /getscores/", ttts.HandleFetchPlayerScores)
	http.HandleFunc("GET /scores/", ttts.HandleScoresView)
}

func (ttts *TicTacToeServer) RegisterAdminHandles() {
	http.HandleFunc("GET /admin/players/", ttts.HandleAdminListPlayers)

	http.HandleFunc("/login/", ttts.HandleLoginView)
	http.HandleFunc("POST /accessc", ttts.HandleAccessControl)
	http.HandleFunc("/adminpage/", ttts.HandleAdminView)
	http.HandleFunc("/fetchdata/", ttts.HandleGetData)
	http.HandleFunc("POST /manualnewplayer/", ttts.HandleNewPlayer)
	http.HandleFunc("POST /regeneratetoken/", ttts.HandleRegenerateToken)
	http.HandleFunc("POST /handleplayeraccess/", ttts.HandleEditPlayerPermissions)
}

func (ttts *TicTacToeServer) AddLobby(lobby *Lobby) {
	ttts.lobbiesLock.Lock()
	// TODO add check if lobby ID exists
	ttts.Lobbies[lobby.LobbyID] = lobby
	ttts.lobbiesLock.Unlock()
}

func (ttts *TicTacToeServer) GetLobby(lobbyId string) (*Lobby, error) {
	return nil, errors.New("Not implemented")
}

func (ttts *TicTacToeServer) Json() ([]byte, error) {
	ttts.lobbiesLock.Lock()
	payload, err := json.Marshal(ttts.Lobbies)
	ttts.lobbiesLock.Unlock()

	return payload, err
}
