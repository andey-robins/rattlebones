package die

import (
	"math/rand"
)

type Die struct {
	Value int
}

func NewDie() *Die {
	return &Die{
		Value: rand.Intn(6) + 1,
	}
}

func NewDieFromValue(val int) *Die {
	return &Die{
		Value: val,
	}
}
