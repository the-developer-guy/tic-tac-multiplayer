package server

import (
	"net/http"
	"sync"
)

type TicTacToeServer struct {
	Lobbies     map[string]Lobby
	lobbiesLock sync.Mutex
}

func NewTicTacToeServer() *TicTacToeServer {
	ttts := TicTacToeServer{
		Lobbies: make(map[string]Lobby),
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
