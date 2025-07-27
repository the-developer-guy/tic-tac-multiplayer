package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TicTacToeMark struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type TicTacToeGrid struct {
	XMarks []TicTacToeMark `json:"x"`
	OMarks []TicTacToeMark `json:"o"`
}

func GetGameGrid(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling get arena")

	writeGridJson(w)
}

func PlaceMark(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling mark placement from player")
}

func writeGridJson(w http.ResponseWriter) {

	grid := TicTacToeGrid{
		XMarks: []TicTacToeMark{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 1,
				Y: 0,
			},
		},
		OMarks: []TicTacToeMark{
			{
				X: 2,
				Y: 2,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grid)
}
