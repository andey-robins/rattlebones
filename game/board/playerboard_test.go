package board

import (
	"rattlebones/game/die"
	"testing"
)

func TestPlayerBoards(t *testing.T) {
	tests := []struct {
		dice   [][]*die.Die
		isFull bool
	}{
		{[][]*die.Die{
			{die.NewDieFromValue(1), die.NewDieFromValue(2), die.NewDieFromValue(3)},
			{die.NewDieFromValue(4), die.NewDieFromValue(5), die.NewDieFromValue(6)},
			{die.NewDieFromValue(1), die.NewDieFromValue(2), die.NewDieFromValue(3)},
		}, true},
		{[][]*die.Die{
			{die.NewDieFromValue(4), die.NewDieFromValue(5), die.NewDieFromValue(6)},
			{die.NewDieFromValue(1), die.NewDieFromValue(2), die.NewDieFromValue(3)},
			{die.NewDieFromValue(4), die.NewDieFromValue(5)},
		}, false},
	}

	for _, test := range tests {
		pb := NewPlayerBoard()
		for i, col := range test.dice {
			for _, d := range col {
				pb.AddDieToColumn(d, i)
			}
		}

		if pb.IsFull() != test.isFull {
			t.Errorf("Expected %t, got %t", test.isFull, pb.IsFull())
		}

		for i, col := range test.dice {
			for j, d := range col {
				if pb.Columns[i].Slots[j].Die.Value != d.Value {
					t.Errorf("Expected %d, got %d", d.Value, pb.Columns[i].Slots[j].Die.Value)
				}
			}
		}

		for i, col := range test.dice {
			for j, d := range col {
				pb.RemoveDieFromColumn(d, i)
				if pb.Columns[i].Slots[j].Die != nil && pb.Columns[i].Slots[j].Die.Value == d.Value {
					t.Errorf("Expected %d to be removed", d.Value)
				}

				for _, s := range pb.Columns[i].Slots {
					if s.Die != nil && s.Die.Value == d.Value {
						t.Errorf("Expected %d to be removed", d.Value)
					}
				}
			}
		}
	}
}
