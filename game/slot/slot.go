package slot

import (
	"rattlebones/game/die"
)

type Slot struct {
	Die *die.Die
}

func NewSlot() *Slot {
	return &Slot{
		Die: nil,
	}
}

func (s *Slot) PlaceDie(d *die.Die) {
	s.Die = d
}
