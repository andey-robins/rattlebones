package board

import (
	"rattlebones/game/column"
	"rattlebones/game/die"
)

type PlayerBoard struct {
	Columns []*column.Column
}

func NewPlayerBoard() *PlayerBoard {
	columns := make([]*column.Column, 3)
	for i := 0; i < 3; i++ {
		columns[i] = column.NewColumn()
	}
	return &PlayerBoard{
		Columns: columns,
	}
}

func (pb *PlayerBoard) IsFull() bool {
	for _, c := range pb.Columns {
		if c.EmptySpaces() > 0 {
			return false
		}
	}
	return true
}

func (pb *PlayerBoard) AddDieToColumn(d *die.Die, col int) {
	pb.Columns[col].AddDie(d)
}

func (pb *PlayerBoard) RemoveDieFromColumn(d *die.Die, col int) {
	pb.Columns[col].RemoveDie(d)
}

func (pb *PlayerBoard) Print() {
	for i := 0; i < 3; i++ {
		for _, c := range pb.Columns {
			if c.Slots[i].Die == nil {
				print("_")
			} else {
				print(c.Slots[i].Die.Value)
			}
			print(" ")
		}
		println()
	}
}

func (pb *PlayerBoard) Score() int {
	score := 0
	for _, c := range pb.Columns {
		score += c.Score()
	}
	return score
}
