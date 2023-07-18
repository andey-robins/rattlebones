package agents

import (
	"rattlebones/game/board"
	"rattlebones/game/die"
)

type Agent interface {
	ChooseMove(*die.Die, *board.GameBoard) int
}
