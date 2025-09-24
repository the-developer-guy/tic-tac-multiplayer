package game

import "errors"

type TicTacToeMark struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type TicTacToeGrid struct {
	Field  [3][3]Mark
	XMarks []TicTacToeMark `json:"x"`
	OMarks []TicTacToeMark `json:"o"`
}

func NewTicTacToeGrid() *TicTacToeGrid {
	tttg := TicTacToeGrid{
		Field: [3][3]Mark{
			{MarkEmpty, MarkEmpty, MarkEmpty},
			{MarkEmpty, MarkEmpty, MarkEmpty},
			{MarkEmpty, MarkEmpty, MarkEmpty},
		},
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
