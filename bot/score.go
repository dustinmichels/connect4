package bot

import (
	"connect4/game"
	"fmt"
)

// Return a score for the given board.
//
//	Positive score means player 1 is winning.
//	Negative score means player 2 is winning.
//	The score is computed based on the number of matching neighbors each piece has.
func Score(board *game.Board) int {

	player1Score := 0
	player2Score := 0

	// factor emphasizes the goodness of having connected pieces
	factor := 10

	for row := 0; row < board.NumRows(); row++ {
		for col := 0; col < board.NumCols(); col++ {
			if board.Get(row, col) == game.EmptySymbol {
				continue
			}
			if board.Get(row, col) == game.Player1Symbol {
				player1Score += countNeighbors(board, row, col, game.Player1Symbol) * factor
			} else {
				player2Score += countNeighbors(board, row, col, game.Player2Symbol) * factor
			}
		}
	}

	fmt.Println("player1Score:", player1Score)
	fmt.Println("player2Score:", player2Score)

	return player1Score - player2Score
}

// Count the number of neighboring pieces (above/below/left/right/diagonal)
// that match the given symbol
func countNeighbors(board *game.Board, row, col int, symbol string) int {
	count := 0

	// check above
	if row > 0 && board.Get(row-1, col) == symbol {
		count++
	}

	// check below
	if row < board.NumRows()-1 && board.Get(row+1, col) == symbol {
		count++
	}

	// check left
	if col > 0 && board.Get(row, col-1) == symbol {
		count++
	}

	// check right
	if col < board.NumCols()-1 && board.Get(row, col+1) == symbol {
		count++
	}

	// check diagonal up and left
	if row > 0 && col > 0 && board.Get(row-1, col-1) == symbol {
		count++
	}

	// check diagonal up and right
	if row > 0 && col < board.NumCols()-1 && board.Get(row-1, col+1) == symbol {
		count++
	}

	// check diagonal down and left
	if row < board.NumRows()-1 && col > 0 && board.Get(row+1, col-1) == symbol {
		count++
	}

	// check diagonal down and right
	if row < board.NumRows()-1 && col < board.NumCols()-1 && board.Get(row+1, col+1) == symbol {
		count++
	}

	fmt.Printf("countNeighbors [%s: %v, %v]: %v\n", symbol, row, col, count)

	return count
}
