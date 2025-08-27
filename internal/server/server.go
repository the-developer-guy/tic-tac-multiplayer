package server

import (
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

func (ttts *TicTacToeServer) RegisterHandles() {
	http.HandleFunc("/grid/{lobbyId}/", ttts.GetGameGrid)
	http.HandleFunc("/place/{lobbyId}/", ttts.PlaceMark)
	http.HandleFunc("/status/{lobbyId}/", ttts.GetLobbyStatus)
	http.HandleFunc("/getlobbies/", ttts.GetActiveLobbies)
	http.HandleFunc("POST /createlobby/", ttts.HandleCreateLobby) // TODO automate lobby creation
}
