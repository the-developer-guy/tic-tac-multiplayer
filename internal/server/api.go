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

	r := writeGridJson()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

func PlaceMark(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling mark placement from player")
}

func writeGridJson() TicTacToeGrid {

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
	return grid
}
