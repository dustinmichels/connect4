package main

import (
	"fmt"
)

// TODO: should this be a pointer?
func MakeMove(b *Board, isPlayer1 bool, col int) (*Board, error) {

	if col < 0 || col >= b.NumCols() {
		return b, fmt.Errorf("invalid column %v", col)
	}

	playerSymbol := Player1Symbol
	if !isPlayer1 {
		playerSymbol = Player2Symbol
	}

	for row := b.NumRows() - 1; row >= 0; row-- {
		if b.grid[row][col] == EmptySymbol {
			b.grid[row][col] = playerSymbol
			return b, nil
		}
	}
	return b, fmt.Errorf("column %v is full", col)
}

func main() {
	board := NewBoard(7, 6)

	b, _ := MakeMove(board, true, 0)

	fmt.Println(b)
	// fmt.Println(board)

}
