package main

import (
	"encoding/json"
	"fmt"
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
		r.ParseForm()
		params := []string{
			r.Form.Get("token"),
			r.Form.Get("cor_x"),
			r.Form.Get("cor_y"),
		}

		if !server.ValidatePOST(w, r, params) {
			return
		}
		new_lobby := server.CreateLobby(r)
		Lobbys = append(Lobbys, new_lobby)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(`{"Lobbyid":` + `"` + new_lobby.LobbyID + `"}`)

		//localhost:8080/LOBBYID
		//Returns the grid for the selected Lobby
		lobbypath := "/" + new_lobby.LobbyID
		http.HandleFunc(lobbypath, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(new_lobby.Grid)
		})

		//place
		http.HandleFunc(lobbypath+"/place", func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			params := []string{
				r.Form.Get("token"),
				r.Form.Get("cor_x"),
				r.Form.Get("cor_y"),
			}
			if !server.ValidatePOST(w, r, params) {
				return
			}
			fmt.Println("Handling Mark Placement")
		})

		http.HandleFunc(lobbypath+"/getstatus", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Handling getstatus")
			//It supposed to return the grid, who's move is the next, and the status which could be either X_won,O_won,Draw or in-game
		})
	})

	http.HandleFunc("/getlobbys", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(Lobbys)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}) //Only for testing
	http.ListenAndServe(":8080", nil)

}
