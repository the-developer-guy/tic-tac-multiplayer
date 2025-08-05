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

	r := GenerateGrid()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

func PlaceMark(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling mark placement from player")
}

func GenerateGrid() TicTacToeGrid {

	grid := TicTacToeGrid{
		XMarks: []TicTacToeMark{},
		OMarks: []TicTacToeMark{},
	}
	return grid
}
