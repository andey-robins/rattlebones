package slot

import (
	"rattlebones/game/die"
	"testing"
)

func TestSlots(t *testing.T) {
	for i := 1; i <= 6; i++ {
		s := NewSlot()
		d := die.NewDieFromValue(i)
		s.PlaceDie(d)
		if s.Die.Value != i {
			t.Errorf("Expected %d, got %d", i, s.Die.Value)
		}
	}
}
