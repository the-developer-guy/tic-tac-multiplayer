package main

import (
	"encoding/json"
	"net/http"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal/server"
)

var Lobbys []server.Lobby

func main() {
	//http://localhost:8080/grid
	http.HandleFunc("/grid", server.GetGameGrid)

	// http://localhost:8080/place
	http.HandleFunc("/place", server.PlaceMark)

	http.HandleFunc("/createlobby", func(w http.ResponseWriter, r *http.Request) {
		/*param := r.URL.Query()
		if len(param) == 0 {

		}*/
		new_lobby := server.CreateLobby(w, r)
		Lobbys = append(Lobbys, new_lobby)
	})

	http.HandleFunc("/getlobbys", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(Lobbys)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		//json.NewEncoder(w).Encode(jsonData)
		w.Write(jsonData)
	}) //Only for testing
	http.ListenAndServe(":8080", nil)

}
