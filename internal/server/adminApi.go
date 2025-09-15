package server

import (
	"net/http"
	"os"
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
	envAdminToken := os.Getenv("ADMIN_TOKEN")

	if adminToken != envAdminToken {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("ok"))
}
