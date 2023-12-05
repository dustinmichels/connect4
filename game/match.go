package game

import (
	"fmt"
	"strconv"
)

type Match struct {
	Board         *Board
	Moves         []int
	Players       []Player
	player1Active bool
}

func NewMatch(players []Player) *Match {
	return &Match{
		Board:         NewBoard(),
		Moves:         []int{},
		Players:       players,
		player1Active: true,
	}
}

func (m *Match) ApplyMove(isPlayer1 bool, col int) error {
	err := m.Board.Update(isPlayer1, col)
	if err != nil {
		return err
	}
	m.Moves = append(m.Moves, col)
	m.player1Active = !m.player1Active

	return nil
}

func getInput() int {

	for {
		fmt.Println("input text:")
		var playerInput string
		_, err := fmt.Scanln(&playerInput)
		if err != nil {
			continue
		}

		// string to int
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

func (m *Match) Start() {

	for {

		// clear screen
		fmt.Print("\033[H\033[2J")
		m.Board.Print()
		fmt.Println(m.Moves)

		// check for winner
		winner, found := m.Board.GetWinner()
		if found {
			fmt.Printf("Player %v wins!\n", winner)
			break
		}

		// print welcome message
		if m.player1Active {
			fmt.Printf("%v's turn\n", m.Players[0].Name)
		} else {
			fmt.Printf("%v's turn\n", m.Players[1].Name)
		}

		// get input
		col := getInput() - 1

		// update board
		err := m.ApplyMove(m.player1Active, col)
		if err != nil {
			panic(err)
		}

	}

}
