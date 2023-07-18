package board

import (
	"rattlebones/game/die"
	"testing"
)

func TestGameBoards(t *testing.T) {
	tests := []struct {
		p1DiceVals [][]int
		p2DiceVals [][]int
		gameover   bool
	}{
		{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{1, 2, 3},
		}, [][]int{
			{4, 5, 6},
			{1, 2, 3},
			{4, 5},
		}, true},
		{[][]int{
			{1, 2, 3},
			{4, 5},
			{1, 2, 3},
		}, [][]int{
			{4, 5, 6},
			{1, 2, 3},
			{4, 5},
		}, false},
	}

	for _, test := range tests {
		gb := NewGameBoard()
		for i, col := range test.p1DiceVals {
			for _, val := range col {
				die := die.NewDieFromValue(val)
				gb.AddDieToColumn(die, i)
			}
		}
		gb.NextTurn()
		for i, col := range test.p2DiceVals {
			for _, val := range col {
				die := die.NewDieFromValue(val)
				gb.AddDieToColumn(die, i)
			}
		}

		if gb.IsGameOver() != test.gameover {
			t.Errorf("Expected %t, got %t", test.gameover, gb.IsGameOver())
		}
	}

	gb := NewGameBoard()
	gb.AddDieToColumn(die.NewDieFromValue(1), 0)
	gb.NextTurn()
	gb.AddDieToColumn(die.NewDieFromValue(1), 0)
	gb.NextTurn()
	if gb.Boards[gb.ActivePlayer].Columns[0].Slots[0].Die != nil {
		t.Errorf("Expected 1 to be removed, got %v", gb.Boards[0].Columns[0].Slots[0].Die)
	}
	gb.NextTurn()
	if gb.Boards[gb.ActivePlayer].Columns[0].Slots[0].Die == nil {
		t.Errorf("Expected 1 to be added, got %v", gb.Boards[0].Columns[0].Slots[0].Die)
	}
}
