package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal/game"
)

func (gs *GameServer) HandlePlayerInfo(w http.ResponseWriter, r *http.Request) {

	playerId := r.PathValue("playerId")
	if playerId == "" {
		http.Error(w, "Missing player ID", http.StatusBadRequest)
		return
	}

	placeholder := "{\"wins\": 0, \"losses\": 0, \"nextGame\": 0}"

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(placeholder))
}

func (gs *GameServer) HandleReadyPlayer(w http.ResponseWriter, r *http.Request) {

	if gs.settings.LocalTest {
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

func (gs *GameServer) HandleGetGameGrid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling game arena getter")

	if gs.settings.LocalTest {
		placeholder := game.NewTicTacToeGrid()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(placeholder)
		return
	}

	lobbyId := r.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}
	_, err := gs.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Nonexistent lobby ID %s", lobbyId), http.StatusBadRequest)
		return
	}

	placeholder := game.NewTicTacToeGrid()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(placeholder)
}

func (gs *GameServer) HandlePlaceMark(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling mark placement from player")

	r.ParseForm()
	token := r.Form.Get("token")
	corX := r.Form.Get("cor_x")
	corY := r.Form.Get("cor_y")
	lobbyId := r.PathValue("lobbyId")
	if corX == "" || corY == "" {
		http.Error(w, "Missing position argument(s)", http.StatusBadRequest)
		return
	}

	if gs.settings.LocalTest {
		if token == "" {
			fmt.Println("Missing player token!")
		}

		if lobbyId == "" {
			fmt.Println("Missing lobby ID!")
			lobbyId = "0"
		}
	} else {
		if token == "" {
			http.Error(w, "Missing Arguments", http.StatusBadRequest)
			return
		}

		if lobbyId == "" {
			http.Error(w, "Missing lobby ID", http.StatusBadRequest)
			return
		}
	}

	lobby, err := gs.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("no lobby with ID %s", lobbyId), http.StatusInternalServerError)
	}

	x, err := strconv.Atoi(corX)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid X coordinate \"%s\"", corX), http.StatusBadRequest)
		return
	}
	y, err := strconv.Atoi(corY)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid Y coordinate \"%s\"", corY), http.StatusBadRequest)
		return
	}

	err = lobby.PlaceMark(x, y, token)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to place mark at %d;%d: %v", x, y, err), http.StatusBadRequest)
		return
	}

	// TODO status 200
}
