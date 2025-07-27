package main

import (
	"net/http"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal/server"
)

func main() {

	//http://localhost:8080/grid
	http.HandleFunc("/grid", server.GetGameGrid)

	// http://localhost:8080/place
	http.HandleFunc("/place", server.PlaceMark)

	http.ListenAndServe(":8080", nil)
}
