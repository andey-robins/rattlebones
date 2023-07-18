package ai

import (
	"math/rand"
	"rattlebones/game/board"
	"rattlebones/game/die"
)

type RandomAgent struct{}

func NewRandomAgent() *RandomAgent {
	return &RandomAgent{}
}

func (ra *RandomAgent) ChooseMove(d *die.Die, gb *board.GameBoard) int {
	for {
		col := rand.Intn(3)
		if gb.Boards[gb.ActivePlayer].Columns[col].EmptySpaces() == 0 {
			continue
		}
		return col
	}
}
