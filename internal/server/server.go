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

func (ttts *TicTacToeServer) GetLobby(lobbyId string) (*Lobby, error) {
	return nil, errors.New("Not implemented")
}

func NewTicTacToeServer() *TicTacToeServer {
	ttts := TicTacToeServer{
		Lobbies: make(map[string]*Lobby),
	}
	return &ttts
}

func (ttts *TicTacToeServer) RegisterApiHandles() {
	http.HandleFunc("/grid/{lobbyId}/", ttts.GetGameGrid)
	http.HandleFunc("/place/{lobbyId}/", ttts.PlaceMark)
	http.HandleFunc("/status/{lobbyId}/", ttts.GetLobbyStatus)
	http.HandleFunc("/getlobbies/", ttts.GetActiveLobbies)
}

func (ttts *TicTacToeServer) RegisterAdminHandles() {
	http.HandleFunc("POST /createlobby/", ttts.HandleCreateLobby) // TODO automate lobby creation
}

func (ttts *TicTacToeServer) AddLobby(lobby *Lobby) {
	ttts.lobbiesLock.Lock()
	// TODO add check if lobby ID exists
	ttts.Lobbies[lobby.LobbyID] = lobby
	ttts.lobbiesLock.Unlock()
}

func (ttts *TicTacToeServer) Json() ([]byte, error) {
	ttts.lobbiesLock.Lock()
	payload, err := json.Marshal(ttts.Lobbies)
	ttts.lobbiesLock.Unlock()

	return payload, err
}
