package board

import (
	"strings"
)

func (b *Board) GetWinner() (string, bool) {

	// check for horizontal wins
	for _, row := range b.grid {
		if winner, ok := checkRow(row); ok {
			return winner, true
		}
	}

	// check for vertical wins
	for col := 0; col < b.NumCols(); col++ {
		colValues := make([]string, b.NumRows())
		for row := 0; row < b.NumRows(); row++ {
			colValues[row] = b.grid[row][col]
		}
		if winner, ok := checkRow(colValues); ok {
			return winner, true
		}
	}

	// check for diagonal wins
	// go along the top row, slanting down and right
	for col := 0; col < b.NumCols(); col++ {
		diagonal := make([]string, b.NumCols())
		for row := 0; row < b.NumRows(); row++ {
			if col+row < b.NumCols() {
				diagonal[row] = b.grid[row][col+row]
			}
		}
		if winner, ok := checkRow(diagonal); ok {
			return winner, true
		}
	}
	// go along the top row, slanting down and left
	for col := 0; col < b.NumCols(); col++ {
		diagonal := make([]string, b.NumCols())
		for row := 0; row < b.NumRows(); row++ {
			if col-row >= 0 {
				diagonal[row] = b.grid[row][col-row]
			}
		}
		if winner, ok := checkRow(diagonal); ok {
			return winner, true
		}
	}

	return "", false
}

func checkRow(row []string) (string, bool) {

	player1Win := strings.Repeat(Player1Symbol, 4)
	player2Win := strings.Repeat(Player2Symbol, 4)

	rowS := strings.Join(row, "")

	if strings.Contains(rowS, player1Win) {
		return Player1Symbol, true
	}

	if strings.Contains(rowS, player2Win) {
		return Player2Symbol, true
	}

	return "", false
}
