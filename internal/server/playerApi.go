package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/the-developer-guy/tic-tac-multiplayer/internal/auth"
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

	pid := r.PathValue("playerId")
	if pid == "" {
		http.Error(w, "Missing player ID", http.StatusBadRequest)
		return
	}
	playerId, err := strconv.ParseInt(pid, 10, 32)
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	r.ParseForm()
	token := r.Form.Get("token")
	if token == "" {
		http.Error(w, "Missing token", http.StatusBadRequest)
		return
	}

	if gs.settings.LocalTest {
		_, err := gs.players.GetPlayer(playerId)
		if err != nil {
			p := auth.NewPlayer("test", token)
			gs.players.AddPlayer(playerId, p)
		}

		l := game.NewLobby(token, "", playerId, 0, time.Now())
		gs.AddLobby(l)

	} else {
		_, err := gs.players.GetAuthenticatedPlayer(playerId, token)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		gs.AddReadyPlayer(playerId)
	}

	lobby, err := gs.GetReadyLobby(playerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if lobby == nil {
		// not scheduled
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{}"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(lobby.ScheduleJson())
}

func (gs *GameServer) HandleGetGameGrid(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling game arena getter")

	lobbyId := r.PathValue("lobbyId")
	if lobbyId == "" {
		http.Error(w, "Missing lobby ID", http.StatusBadRequest)
		return
	}
	lobby, err := gs.GetLobby(lobbyId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Nonexistent lobby ID %s", lobbyId), http.StatusBadRequest)
		return
	}

	j, err := lobby.GridJson()
	if err != nil {
		http.Error(w, "failed to parse game arena", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
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
			for lid := range gs.ActiveTournamentLobbies {
				lobbyId = lid
				break
			}
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
