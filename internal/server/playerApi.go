package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (ttts *TicTacToeServer) GetGameGrid(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling get arena")

	lobbyId := req.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}
	_, err := ttts.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Nonexistent lobby ID %s", lobbyId), http.StatusBadRequest)
		return
	}

	r := ttts.GenerateGrid()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

func (ttts *TicTacToeServer) PlaceMark(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling mark placement from player")

	lobbyId := req.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}
	_, err := ttts.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Nonexistent lobby ID %s", lobbyId), http.StatusBadRequest)
		return
	}
}

func (ttts *TicTacToeServer) GetLobbyStatus(w http.ResponseWriter, req *http.Request) {
	fmt.Println("getting lobby status")

	lobbyId := req.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}
	_, err := ttts.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Nonexistent lobby ID %s", lobbyId), http.StatusBadRequest)
		return
	}
}

func HandlePlaceInLobby(w http.ResponseWriter, r *http.Request, lobbyPath string) {
	r.ParseForm()
	token := r.Form.Get("token")
	corX := r.Form.Get("cor_x")
	corY := r.Form.Get("cor_y")
	if token == "" || corX == "" || corY == "" {
		http.Error(w, "Missing Arguments", http.StatusBadRequest)
		return
	}
	fmt.Println("Handling Mark Placement")
}

func HandleGetStatusInLobby(w http.ResponseWriter, r *http.Request, lobbyPath string) {
	fmt.Println("Handling getstatus")
	//lobbyPath is going to come in handy when it comes to actually handling the request.
}

func (ttts *TicTacToeServer) GetActiveLobbies(w http.ResponseWriter, r *http.Request) {

	lobbiesJson, err := ttts.Json()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(lobbiesJson)
}
