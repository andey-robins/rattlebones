package column

import (
	"rattlebones/game/die"
	"testing"
)

func TestColumns(t *testing.T) {
	tests := []struct {
		dice  []*die.Die
		score int
	}{
		{[]*die.Die{die.NewDieFromValue(1), die.NewDieFromValue(2), die.NewDieFromValue(3)}, 6},
		{[]*die.Die{die.NewDieFromValue(4), die.NewDieFromValue(5), die.NewDieFromValue(6)}, 15},
		{[]*die.Die{die.NewDieFromValue(1)}, 1},
		{[]*die.Die{die.NewDieFromValue(2), die.NewDieFromValue(3)}, 5},
	}

	for _, test := range tests {
		c := NewColumn()
		for i, d := range test.dice {
			c.AddDie(d)
			if c.Slots[i].Die.Value != d.Value {
				t.Errorf("Expected %d, got %d", d.Value, c.Slots[i].Die.Value)
			}
		}

		if c.EmptySpaces() != 3-len(test.dice) {
			t.Errorf("Expected %d, got %d", 3-len(test.dice), c.EmptySpaces())
		}

		if c.Score() != test.score {
			t.Errorf("Expected %d, got %d", test.score, c.Score())
		}

		deleted := 0

		for _, d := range test.dice {
			c.RemoveDie(d)
			deleted++

			for _, s := range c.Slots {
				if s.Die != nil && s.Die.Value == d.Value {
					t.Errorf("Expected %d to be removed", d.Value)
				}
			}

			if c.EmptySpaces() != 3-len(test.dice)+deleted {
				t.Errorf("Expected %d, got %d", 3-len(test.dice)+deleted, c.EmptySpaces())
			}

			if c.Slots[0].Die == nil && c.Slots[1].Die != nil {
				t.Errorf("Expected %d to be moved down", c.Slots[1].Die.Value)
			}
		}
	}
}
