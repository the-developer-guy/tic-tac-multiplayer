package server

import "sync"

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
