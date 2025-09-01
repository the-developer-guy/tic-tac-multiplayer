package server

import (
	"encoding/json"
	"net/http"
)

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
	l := CreateLobbyFromRequest(r)
	ttts.AddLobby(l)

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"Lobbyid": l.LobbyID}
	json.NewEncoder(w).Encode(response)
}

func CreateLobbyFromRequest(req *http.Request) *Lobby {
	req.ParseForm()
	token1 := req.Form.Get("token")
	token2 := req.Form.Get("token2")

	return NewLobby(token1, token2)
}
