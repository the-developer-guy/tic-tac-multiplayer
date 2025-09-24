package game

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

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

func (tttg *TicTacToeGrid) PlaceMark(x, y int, mark Mark) error {

	if tttg.Field[x][y] != MarkEmpty {
		return fmt.Errorf("%d;%d position is already filled", x, y)
	}

	m := TicTacToeMark{x, y}

	switch mark {
	case MarkX:
		tttg.XMarks = append(tttg.XMarks, m)
	case MarkO:
		tttg.OMarks = append(tttg.OMarks, m)
	case MarkEmpty:
		return errors.New("can't place empty mark")
	default:
		return errors.New("invalid mark")
	}

	tttg.Field[x][y] = mark

	return nil
}

func (tttg *TicTacToeGrid) PlaceRandomMark(mark Mark) error {

	freePlaces := [9]TicTacToeMark{}
	freeCount := 0

	for x, row := range tttg.Field {
		for y, cell := range row {
			if cell == MarkEmpty {
				freePlaces[freeCount].X = x
				freePlaces[freeCount].Y = y
				freeCount++
			}
		}
	}
	if freeCount == 0 {
		return errors.New("game field is full")
	}

	r := rand.IntN(freeCount)
	m := freePlaces[r]

	return tttg.PlaceMark(m.X, m.Y, mark)
}
