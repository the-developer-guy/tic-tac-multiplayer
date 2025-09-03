package main

import (
	"net/http"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal/server"
)

func main() {

	s := server.NewTicTacToeServer()
	s.RegisterApiHandles()
	s.RegisterAdminHandles()

	http.ListenAndServe(":8080", nil)
}
