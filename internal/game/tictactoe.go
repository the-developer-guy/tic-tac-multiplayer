package game

import "errors"

type TicTacToeMark struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type TicTacToeGrid struct {
	XMarks []TicTacToeMark `json:"x"`
	OMarks []TicTacToeMark `json:"o"`
}

func NewTicTacToeGrid() *TicTacToeGrid {
	tttg := TicTacToeGrid{
		XMarks: []TicTacToeMark{},
		OMarks: []TicTacToeMark{},
	}

	return &tttg
}

func (tttg *TicTacToeGrid) PlaceXMark(x, y int) error {
	return errors.New("not implemented")
}

func (tttg *TicTacToeGrid) PlaceOMark(x, y int) error {
	return errors.New("not implemented")
}

func (tttg *TicTacToeGrid) PlaceRandomOMark() error {
	return errors.New("not implemented")
}
