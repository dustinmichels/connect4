package game

import (
	"strings"
)

// Returns the winner of the game as int (1 or 2) if found
func (b *Board) GetWinner() (winner int, found bool) {

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

	// check for diagonal wins, slanting down and right
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

	// check for diagonal wins, slanting down and left
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

	return 0, false

}

func checkRow(row []string) (int, bool) {
	player1Win := strings.Repeat(Player1Symbol, 4)
	player2Win := strings.Repeat(Player2Symbol, 4)
	rowS := strings.Join(row, "")

	if strings.Contains(rowS, player1Win) {
		return 1, true
	}

	if strings.Contains(rowS, player2Win) {
		return 2, true
	}

	return 0, false
}
