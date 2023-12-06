package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/enescakir/emoji"
)

const (
	EmptySymbol   = "-"
	Player1Symbol = "X"
	Player2Symbol = "O"
	boardWidth    = 7
	boardHeight   = 6
)

type Board struct {
	grid [][]string
}

func NewBoard() *Board {
	grid := make([][]string, boardHeight)
	for i := range grid {
		grid[i] = make([]string, boardWidth)
		for j := range grid[i] {
			grid[i][j] = EmptySymbol
		}
	}
	return &Board{grid}
}

func (b *Board) NumCols() int {
	return len(b.grid[0])
}

func (b *Board) NumRows() int {
	return len(b.grid)
}

func (b *Board) Get(row, col int) string {
	return b.grid[row][col]
}

// ---------- Applying moves ----------

// Add the appropriate symbol to the first empty row in the given column
// Returns an error if the column is full or invalid
func (b *Board) Update(isPlayer1 bool, col int) error {

	if col < 0 || col >= b.NumCols() {
		return fmt.Errorf("invalid column %v", col)
	}

	playerSymbol := Player1Symbol
	if !isPlayer1 {
		playerSymbol = Player2Symbol
	}

	for row := b.NumRows() - 1; row >= 0; row-- {
		if b.grid[row][col] == EmptySymbol {
			b.grid[row][col] = playerSymbol
			return nil
		}
	}
	return fmt.Errorf("column %v is full", col)
}

// Update board with a series of columns chosen by alternating players
// Eg, board.ApplyMoves([]int{1, 2, 3, 4, 5, 6, 7})
func (b *Board) ApplyMoves(moves []int) error {
	for i, col := range moves {
		if err := b.Update(i%2 == 0, col); err != nil {
			return err
		}
	}
	return nil
}

// ---------- Check status ----------

// Check if top row is all full
func (b *Board) IsFull() bool {
	for _, col := range b.grid[0] {
		if col == EmptySymbol {
			return false
		}
	}
	return true
}

// Returns the winner of the game as int (1 or 2) if found
func (b *Board) GetWinner() (winner int, found bool) {

	// check for horizontal wins
	for _, row := range b.grid {
		if winner, found := checkRow(row); found {
			return winner, true
		}
	}

	// check for vertical wins
	for col := 0; col < b.NumCols(); col++ {
		colValues := make([]string, b.NumRows())
		for row := 0; row < b.NumRows(); row++ {
			colValues[row] = b.grid[row][col]
		}
		if winner, found := checkRow(colValues); found {
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
		if winner, found := checkRow(diagonal); found {
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
		if winner, found := checkRow(diagonal); found {
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

// ---------- Display ----------

func (b *Board) Print() {
	fmt.Println(b.String())
}

func (b *Board) String() string {

	displayTable := map[string]emoji.Emoji{
		EmptySymbol:   emoji.WhiteCircle,
		Player1Symbol: emoji.RedCircle,
		Player2Symbol: emoji.YellowCircle,
		"1":           emoji.Keycap1,
		"2":           emoji.Keycap2,
		"3":           emoji.Keycap3,
		"4":           emoji.Keycap4,
		"5":           emoji.Keycap5,
		"6":           emoji.Keycap6,
		"7":           emoji.Keycap7,
		"8":           emoji.Keycap8,
		"9":           emoji.Keycap9,
		"10":          emoji.Keycap10,
	}

	s := ""

	// print header
	for i := 0; i < len(b.grid[0]); i++ {
		s += fmt.Sprintf("%v  ", displayTable[strconv.Itoa(i+1)])
	}
	s += "\n"

	// print board
	for _, row := range b.grid {
		for _, col := range row {
			s += fmt.Sprintf("%v ", displayTable[col])
		}
		s += "\n"
	}
	return s
}
