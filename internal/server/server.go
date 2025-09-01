package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
)

type TicTacToeServer struct {
	Lobbies     map[string]*Lobby
	lobbiesLock sync.Mutex
}

func NewTicTacToeServer() *TicTacToeServer {
	ttts := TicTacToeServer{
		Lobbies: make(map[string]*Lobby),
	}
	return &ttts
}

func (ttts *TicTacToeServer) RegisterApiHandles() {
	http.HandleFunc("GET /playerinfo/{playerId}", ttts.HandlePlayerInfo)
	http.HandleFunc("GET /ready/{playerId}/", ttts.HandleReadyPlayer)

	http.HandleFunc("GET /getgrid/{lobbyId}/", ttts.HandleGetGameGrid)
	http.HandleFunc("POST /place/{lobbyId}/", ttts.HandlePlaceMark)
}

func (ttts *TicTacToeServer) RegisterAdminHandles() {
	http.HandleFunc("GET /admin/players/", ttts.HandleAdminListPlayers)
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
