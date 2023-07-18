package column

import (
	"rattlebones/game/die"
	"rattlebones/game/slot"
)

type Column struct {
	Slots []*slot.Slot
}

func NewColumn() *Column {
	slots := make([]*slot.Slot, 3)
	for i := 0; i < 3; i++ {
		slots[i] = slot.NewSlot()
	}

	return &Column{
		Slots: slots,
	}
}

func (c *Column) AddDie(d *die.Die) {
	for _, s := range c.Slots {
		if s.Die == nil {
			s.PlaceDie(d)
			return
		}
	}
}

func (c *Column) RemoveDie(d *die.Die) {
	for i, s := range c.Slots {
		if s.Die != nil && s.Die.Value == d.Value {
			c.Slots[i] = slot.NewSlot()
		}
	}
	for i := 0; i < 2; i++ {
		if c.Slots[i].Die == nil && c.Slots[i+1].Die != nil {
			c.Slots[i] = c.Slots[i+1]
			c.Slots[i+1] = slot.NewSlot()
		}
	}
}

func (c *Column) EmptySpaces() int {
	emptySpaces := 0
	for _, s := range c.Slots {
		if s.Die == nil {
			emptySpaces++
		}
	}
	return emptySpaces
}

func (c *Column) Score() int {
	counts := make(map[int]int)
	for _, s := range c.Slots {
		if s.Die != nil {
			counts[s.Die.Value]++
		}
	}

	score := 0
	for k, v := range counts {
		if v == 1 {
			score += k
		} else if v == 2 {
			score += k * 4
		} else if v == 3 {
			score += k * 9
		}
	}
	return score
}
