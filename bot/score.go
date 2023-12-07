package bot

import (
	"connect4/game"
	"fmt"
	"strings"
)

const debug = false

// Return a score for the given board.
//
//	Positive score means player 1 is winning.
//	Negative score means player 2 is winning.
//	The score is computed based on the number of matching neighbors each piece has.
func Score(b *game.Board) int {

	player1Score := 0
	player2Score := 0

	// check how many connected in each row
	for _, row := range b.Grid {
		player1Score += scoreRow(row, game.Player1Symbol)
		player2Score += scoreRow(row, game.Player2Symbol)
		if debug {
			fmt.Println("row   scores: ", player1Score, player2Score)
		}
	}

	// check for vertical wins
	for col := 0; col < b.NumCols(); col++ {
		colValues := make([]string, b.NumRows())
		for row := 0; row < b.NumRows(); row++ {
			colValues[row] = b.Grid[row][col]
		}
		player1Score += scoreRow(colValues, game.Player1Symbol)
		player2Score += scoreRow(colValues, game.Player2Symbol)
		if debug {
			fmt.Println("col   scores: ", player1Score, player2Score)
		}
	}

	// check for diagonal wins, slanting down and right
	for col := 0; col < b.NumCols(); col++ {
		diagonal := make([]string, b.NumCols())
		for row := 0; row < b.NumRows(); row++ {
			if col+row < b.NumCols() {
				diagonal[row] = b.Grid[row][col+row]
			}
		}
		player1Score += scoreRow(diagonal, game.Player1Symbol)
		player2Score += scoreRow(diagonal, game.Player2Symbol)
		if debug {
			fmt.Println("diag1 scores: ", player1Score, player2Score)
		}
	}

	// check for diagonal wins, slanting down and left
	for col := 0; col < b.NumCols(); col++ {
		diagonal := make([]string, b.NumCols())
		for row := 0; row < b.NumRows(); row++ {
			if col-row >= 0 {
				diagonal[row] = b.Grid[row][col-row]
			}
		}
		player1Score += scoreRow(diagonal, game.Player1Symbol)
		player2Score += scoreRow(diagonal, game.Player2Symbol)
		if debug {
			fmt.Println("diag2 scores: ", player1Score, player2Score)
		}
	}

	return player1Score - player2Score

}

func scoreRow(row []string, symbol string) int {
	rowS := strings.Join(row, "")
	switch {
	case strings.Contains(rowS, strings.Repeat(symbol, 4)):
		return 100
	case strings.Contains(rowS, strings.Repeat(symbol, 3)):
		return 10
	case strings.Contains(rowS, strings.Repeat(symbol, 2)):
		return 1
	default:
		return 0
	}
}
