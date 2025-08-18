package server

import "sync"

type TicTacToeServer struct {
	Lobbies     []Lobby
	lobbiesLock sync.Mutex
}

func NewTicTacToeServer() *TicTacToeServer {
	ttts := TicTacToeServer{}

	return &ttts
}
