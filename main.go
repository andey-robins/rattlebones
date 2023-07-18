package main

import (
	"fmt"
	"rattlebones/agents"
	"rattlebones/agents/ai"
	"rattlebones/agents/human"
	"rattlebones/game/board"
	"rattlebones/game/die"
)

func main() {
	PlayVersusRandom()
}

func PlayVersusRandom() {
	game := board.NewGameBoard()

	agents := [2]struct {
		agent agents.Agent
		name  string
	}{
		{human.NewHumanAgent(), "Player"},
		{ai.NewRandomAgent(), "AI"},
	}

	for !game.IsGameOver() {
		activeAgent := agents[game.ActivePlayer].agent
		d := die.NewDie()
		col := activeAgent.ChooseMove(d, game)
		game.AddDieToColumn(d, col)
		game.NextTurn()
	}

	fmt.Printf("%s scores %d points.\n", agents[game.InactivePlayer].name, game.Boards[game.InactivePlayer].Score())
	fmt.Printf("%s scores %d points.\n", agents[game.ActivePlayer].name, game.Boards[game.ActivePlayer].Score())

	game.PrintWithNames(agents[game.InactivePlayer].name, agents[game.ActivePlayer].name)
}
