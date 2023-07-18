package human

import (
	"fmt"
	"rattlebones/game/board"
	"rattlebones/game/die"
)

type HumanAgent struct{}

func NewHumanAgent() *HumanAgent {
	return &HumanAgent{}
}

func (ha *HumanAgent) ChooseMove(d *die.Die, gb *board.GameBoard) int {
	gb.Print()

	var col int
	for {
		fmt.Printf("You have rolled a %d.\n", d.Value)
		println("Which column would you like to place your die in?")
		_, err := fmt.Scanf("%d", &col)
		if err != nil {
			println("Please enter a number.")
			continue
		}
		if col < 0 || col > 2 {
			println("Please enter a number between 0 and 2.")
			continue
		}
		if gb.Boards[gb.ActivePlayer].Columns[col].EmptySpaces() == 0 {
			println("That column is full.")
			continue
		}
		break
	}

	return col
}
