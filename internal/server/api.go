package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (ttts *TicTacToeServer) GetGameGrid(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling get arena")

	r := ttts.GenerateGrid()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

func (ttts *TicTacToeServer) PlaceMark(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling mark placement from player")
}

func (ttts *TicTacToeServer) GenerateGrid() TicTacToeGrid {

	grid := TicTacToeGrid{
		XMarks: []TicTacToeMark{},
		OMarks: []TicTacToeMark{},
	}
	return grid
}

func (ttts *TicTacToeServer) HandleCreateLobby(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	adminToken := r.Form.Get("atoken")
	if adminToken != "admin" { //TODO: Must be replaced in the future
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	newLobby := CreateLobbyFromRequest(r, len(ttts.Lobbies))

	ttts.lobbiesLock.Lock()
	ttts.Lobbies[newLobby.LobbyID] = newLobby
	ttts.lobbiesLock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"Lobbyid": newLobby.LobbyID}
	json.NewEncoder(w).Encode(response)

	lobbyPath := fmt.Sprintf("/%s", newLobby.LobbyID)

	HandlePlaceInLobby(w, r, lobbyPath)     //lobbyID/place
	HandleGetStatusInLobby(w, r, lobbyPath) //lobbyID/getstatus
}

func HandlePlaceInLobby(w http.ResponseWriter, r *http.Request, lobbyPath string) {
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
}

func HandleGetStatusInLobby(w http.ResponseWriter, r *http.Request, lobbyPath string) {
	statusPath := fmt.Sprintf("%s/getstatus", lobbyPath)
	http.HandleFunc(statusPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling getstatus")
		//It's supposed to return the grid, who's move is the next, and the status which could be either X_won,O_won,Draw or in-game
	})
}

func (ttts *TicTacToeServer) GetActiveLobbies(w http.ResponseWriter, r *http.Request) {
	ttts.lobbiesLock.Lock()
	jsonData, err := json.Marshal(ttts.Lobbies)
	ttts.lobbiesLock.Unlock()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
