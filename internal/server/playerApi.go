package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (ttts *TicTacToeServer) HandlePlayerInfo(w http.ResponseWriter, r *http.Request) {

	playerId := r.PathValue("playerId")
	if playerId == "" {
		http.Error(w, "Missing player ID", http.StatusBadRequest)
		return
	}

	placeholder := "{\"wins\": 0, \"losses\": 0, \"nextGame\": 0}"

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(placeholder))
}

func (ttts *TicTacToeServer) HandleReadyPlayer(w http.ResponseWriter, r *http.Request) {

	if ttts.settings.Standalone {
		placeholder := "{\"lobbyId\": \"0\", \"nextGame\": 0}"

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(placeholder))
		return
	}

	playerId := r.PathValue("playerId")
	if playerId == "" {
		http.Error(w, "Missing player ID", http.StatusBadRequest)
		return
	}

	r.ParseForm()
	token := r.Form.Get("token")
	if token == "" {
		http.Error(w, "Missing token", http.StatusBadRequest)
		return
	}

	placeholder := "{\"lobbyId\": \"\", \"nextGame\": 0}"

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(placeholder))
}

func (ttts *TicTacToeServer) HandleGetGameGrid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling game arena getter")

	if ttts.settings.Standalone {
		placeholder := ttts.GenerateGrid()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(placeholder)
		return
	}

	lobbyId := r.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}
	_, err := ttts.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Nonexistent lobby ID %s", lobbyId), http.StatusBadRequest)
		return
	}

	placeholder := ttts.GenerateGrid()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(placeholder)
}

func (ttts *TicTacToeServer) HandlePlaceMark(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling mark placement from player")

	if ttts.settings.Standalone {
		r.ParseForm()
		token := r.Form.Get("token")
		corX := r.Form.Get("cor_x")
		corY := r.Form.Get("cor_y")
		if token == "" || corX == "" || corY == "" {
			http.Error(w, "Missing Arguments", http.StatusBadRequest)
			return
		}
	}

	lobbyId := r.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}

	r.ParseForm()
	token := r.Form.Get("token")
	corX := r.Form.Get("cor_x")
	corY := r.Form.Get("cor_y")
	if token == "" || corX == "" || corY == "" {
		http.Error(w, "Missing Arguments", http.StatusBadRequest)
		return
	}
}
