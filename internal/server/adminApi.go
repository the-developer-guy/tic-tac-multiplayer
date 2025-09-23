package server

import (
	"net/http"
)

func (ttts *TicTacToeServer) GenerateGrid() TicTacToeGrid {

	grid := TicTacToeGrid{
		XMarks: []TicTacToeMark{},
		OMarks: []TicTacToeMark{},
	}
	return grid
}

func (ttts *TicTacToeServer) HandleAdminListPlayers(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	adminToken := r.Form.Get("admintoken")

	if adminToken != ttts.settings.AdminToken {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("ok"))
}
