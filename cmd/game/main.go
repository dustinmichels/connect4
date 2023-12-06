package main

import (
	"connect4/game"
	"fmt"
	"strconv"
)

func GameLoop() {
	g := game.NewGame()
	m := g.Match

	for {
		// reset screen and print board
		fmt.Print("\033[H\033[2J")
		m.Board.Print()
		fmt.Println(m.Moves)

		// check for winner
		winner, found := m.Board.GetWinner()
		if found {
			fmt.Printf("Player %v wins!\n", winner)
			break
		}

		// if no winner, proceed
		if m.IsPlayer1Active() {
			fmt.Printf("%v's turn\n", m.Players[0].Name)
		} else {
			fmt.Printf("%v's turn\n", m.Players[1].Name)
		}

		// get input & update board
		for {
			col := getInput() - 1
			err := m.ApplyMove(col)
			if err == nil {
				break
			}
			fmt.Println("Invalid move")
		}

	}

}

// Infinite loop until player
func getInput() int {
	for {
		fmt.Println("input column of choice:")
		var playerInput string
		_, err := fmt.Scanln(&playerInput)
		if err != nil {
			continue
		}
		colEntry, err := strconv.Atoi(playerInput)
		if err != nil {
			continue
		}
		if colEntry < 1 || colEntry > 7 {
			continue
		}
		return colEntry
	}
}

func main() {
	GameLoop()
}
