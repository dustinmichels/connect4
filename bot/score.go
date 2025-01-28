package bot

import (
	"connect4/game"
)

// const debug = true

func Score(b *game.Board) int {
	return 0
}

// Return a score for the given board.
//
//	Positive score means player1 is winning.
//	Negative score means player2 is winning.
//	The score is computed based on the number of matching neighbors each piece has.
//		- 100 points for 4 in a row
//		- 10 points for 3 in a row
//		- 1 point for 2 in a row
// func Score(b *game.Board) int {

// 	player1Score := 0
// 	player2Score := 0

// 	// check how many connected in each row
// 	for i, row := range b.Grid {
// 		player1Score += scoreRow(row, game.Player1Symbol)
// 		player2Score += scoreRow(row, game.Player2Symbol)
// 		if debug {
// 			fmt.Printf("row   %d: %v: %d %d\n", i, row, player1Score, player2Score)
// 		}
// 	}

// 	// score columns
// 	for col := 0; col < b.NumCols(); col++ {
// 		colValues := make([]string, b.NumRows())
// 		for row := 0; row < b.NumRows(); row++ {
// 			colValues[row] = b.Grid[row][col]
// 		}
// 		player1Score += scoreRow(colValues, game.Player1Symbol)
// 		player2Score += scoreRow(colValues, game.Player2Symbol)
// 		if debug {
// 			fmt.Printf("col   %d: %v  : %d %d\n", col, colValues, player1Score, player2Score)
// 		}
// 	}

// 	// score diagonals, slanting down and right
// 	for col := 0; col < b.NumCols(); col++ {
// 		diagonal := make([]string, b.NumCols())
// 		for row := 0; row < b.NumRows(); row++ {
// 			if col+row < b.NumCols() {
// 				diagonal[row] = b.Grid[row][col+row]
// 			}
// 		}
// 		player1Score += scoreRow(diagonal, game.Player1Symbol)
// 		player2Score += scoreRow(diagonal, game.Player2Symbol)

// 		if debug {
// 			padding := strings.Repeat(" ", col)
// 			if col == 0 {
// 				padding = " "
// 			}
// 			fmt.Printf("diag1 %d: %v%v: %d %d\n", col, diagonal, padding, player1Score, player2Score)
// 		}
// 	}

// 	// score diagonals, slanting down and left
// 	for col := 0; col < b.NumCols(); col++ {
// 		diagonal := make([]string, b.NumCols())
// 		for row := 0; row < b.NumRows(); row++ {
// 			if col-row >= 0 {
// 				diagonal[row] = b.Grid[row][col-row]
// 			}
// 		}
// 		player1Score += scoreRow(diagonal, game.Player1Symbol)
// 		player2Score += scoreRow(diagonal, game.Player2Symbol)
// 		if debug {
// 			padding := strings.Repeat(" ", b.NumCols()-col-1)
// 			if col == b.NumCols()-1 {
// 				padding = " "
// 			}
// 			fmt.Printf("diag2 %d: %v%v: %d %d\n", col, diagonal, padding, player1Score, player2Score)
// 		}
// 	}

// 	return player1Score - player2Score

// }

// func scoreRow(row []string, symbol uint8) int {
// 	rowS := strings.Join(row, "")
// 	switch {
// 	case strings.Contains(rowS, strings.Repeat(symbol, 4)):
// 		return 100
// 	case strings.Contains(rowS, strings.Repeat(symbol, 3)):
// 		return 10
// 	case strings.Contains(rowS, strings.Repeat(symbol, 2)):
// 		return 1
// 	default:
// 		return 0
// 	}
// }
