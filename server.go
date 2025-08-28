package main

import (
	"net/http"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal/server"
)

func main() {

	s := server.NewTicTacToeServer()
	s.RegisterHandles()

	http.ListenAndServe(":8080", nil)
}
