package main

import (
	"net/http"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal"
	"github.com/the-developer-guy/tic-tac-multiplayer/internal/server"
)

func main() {

	config, err := internal.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	s := server.NewTicTacToeServer(config)
	s.RegisterApiHandles()
	s.RegisterAdminHandles()

	http.ListenAndServe(":8080", nil)
}
