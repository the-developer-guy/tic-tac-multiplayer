package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal/server"
)

var (
	Lobbies     []server.Lobby
	lobbiesLock sync.Mutex
)

func main() {
	http.HandleFunc("/grid", server.GetGameGrid)
	http.HandleFunc("/place", server.PlaceMark)
	http.HandleFunc("/getlobbies", handleGetLobbies)
	http.HandleFunc("POST /createlobby", handleCreateLobby)

	http.ListenAndServe(":8080", nil)
}

func handleCreateLobby(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newLobby := server.ReturnLobby(r, len(Lobbies))

	lobbiesLock.Lock()
	Lobbies = append(Lobbies, newLobby)
	lobbiesLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"Lobbyid": newLobby.LobbyID}
	json.NewEncoder(w).Encode(response)

	lobbyPath := fmt.Sprintf("/%s", newLobby.LobbyID)
	http.HandleFunc(lobbyPath, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newLobby.Grid)
	})

	placePath := fmt.Sprintf("POST %s/place", lobbyPath)
	http.HandleFunc(placePath, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		token := r.Form.Get("token")
		corX := r.Form.Get("cor_x")
		corY := r.Form.Get("cor_y")
		if token == "" || corX == "" || corY == "" {
			http.Error(w, "Missing Arguments", http.StatusBadRequest)
			return
		}
		fmt.Println("Handling Mark Placement")
	})

	joinLobbyPath := fmt.Sprintf("POST %s/join", lobbyPath)
	http.HandleFunc(joinLobbyPath, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		token := r.Form.Get("token")
		if newLobby.Players[1].Token != "" || len(token) == 0 {
			http.Error(w, "Lobby is already occupied or wrong token", http.StatusBadRequest)
			return
		}
		newLobby.Players[1].Token = token

		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(newLobby)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	})

	statusPath := fmt.Sprintf("%s/getstatus", lobbyPath)
	http.HandleFunc(statusPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling getstatus")
		//It's supposed to return the grid, who's move is the next, and the status which could be either X_won,O_won,Draw or in-game
	})
}

func handleGetLobbies(w http.ResponseWriter, r *http.Request) {
	lobbiesLock.Lock()
	jsonData, err := json.Marshal(Lobbies)
	lobbiesLock.Unlock()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
