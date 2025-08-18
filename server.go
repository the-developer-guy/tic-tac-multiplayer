package main

import (
	"net/http"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal/server"
)

func main() {

	s := server.NewTicTacToeServer()

	http.HandleFunc("/grid", s.GetGameGrid)
	http.HandleFunc("/place", s.PlaceMark)
	http.HandleFunc("/getlobbies", s.HandleGetLobbies)
	http.HandleFunc("POST /createlobby", s.HandleCreateLobby)

	http.ListenAndServe(":8080", nil)
}
