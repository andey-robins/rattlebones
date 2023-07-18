package board

import (
	"fmt"
	"rattlebones/game/die"
)

type GameBoard struct {
	Boards         []*PlayerBoard
	ActivePlayer   int
	InactivePlayer int
}

func NewGameBoard() *GameBoard {
	boards := make([]*PlayerBoard, 2)
	boards[0] = NewPlayerBoard()
	boards[1] = NewPlayerBoard()
	return &GameBoard{
		Boards:         boards,
		ActivePlayer:   0,
		InactivePlayer: 1,
	}
}

func (gb *GameBoard) NextTurn() {
	gb.ActivePlayer, gb.InactivePlayer = gb.InactivePlayer, gb.ActivePlayer
}

func (gb *GameBoard) AddDieToColumn(d *die.Die, col int) {
	gb.Boards[gb.ActivePlayer].AddDieToColumn(d, col)
	gb.Boards[gb.InactivePlayer].RemoveDieFromColumn(d, col)
}

func (gb *GameBoard) IsGameOver() bool {
	return gb.Boards[gb.ActivePlayer].IsFull() || gb.Boards[gb.InactivePlayer].IsFull()
}

func (gb *GameBoard) Print() {
	fmt.Println("\nYou:")
	gb.Boards[gb.ActivePlayer].Print()
	fmt.Println("\nYour Opponent:")
	gb.Boards[gb.InactivePlayer].Print()
}

func (gb *GameBoard) PrintWithNames(activeName, inactiveName string) {
	fmt.Printf("\n%s:\n", activeName)
	gb.Boards[gb.ActivePlayer].Print()
	fmt.Printf("\n%s:\n", inactiveName)
	gb.Boards[gb.InactivePlayer].Print()
}
