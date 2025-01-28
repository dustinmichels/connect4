package game

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/enescakir/emoji"
)

const (
	EmptySymbol   uint8 = 0
	Player1Symbol uint8 = 1
	Player2Symbol uint8 = 2
	boardWidth          = 7
	boardHeight         = 6
	winningLength       = 4
)

type Board struct {
	Grid [][]uint8
}

func NewBoard() *Board {
	grid := make([][]uint8, boardHeight)
	for i := range grid {
		grid[i] = make([]uint8, boardWidth)
	}
	return &Board{grid}
}

func (b *Board) NumCols() int {
	return len(b.Grid[0])
}

func (b *Board) NumRows() int {
	return len(b.Grid)
}

func (b *Board) Get(row, col int) uint8 {
	return b.Grid[row][col]
}

// ---------- Applying moves ----------

// Add the appropriate symbol to the first empty row in the given column
// Ie, "drop" a piece in the column
// Returns an error if the column is full or invalid
func (b *Board) Update(isPlayer1 bool, col int) error {

	if col < 0 || col >= b.NumCols() {
		return fmt.Errorf("invalid column %v", col)
	}

	// set correct player symbol
	var playerSymbol uint8 = Player2Symbol
	if isPlayer1 {
		playerSymbol = Player1Symbol
	}

	for row := b.NumRows() - 1; row >= 0; row-- {
		if b.Grid[row][col] == 0 {
			b.Grid[row][col] = playerSymbol
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

// Get possible moves for the next player
func (b *Board) GetPossibleMoves() []int {
	moves := []int{}
	for col := 0; col < b.NumCols(); col++ {
		if b.Grid[0][col] == EmptySymbol {
			moves = append(moves, col)
		}
	}
	return moves
}

// Check if top row is all full
func (b *Board) IsFull() bool {
	for _, col := range b.Grid[0] {
		if col == EmptySymbol {
			return false
		}
	}
	return true
}

// func (b *Board) GetWinner() (winner int, found bool) {
// 	return 0, false
// }

// CheckWinner checks if there is a winner on the board.
// Returns the symbol of the winner (Player1Symbol or Player2Symbol) or EmptySymbol if no winner.
func (b *Board) CheckWinner() uint8 {
	return checkWinner(b)
}

// ---------- Display ----------

func (b *Board) Print() {
	fmt.Println(b.String())
}

func (b *Board) String() string {

	displayTable := map[string]emoji.Emoji{
		"1":  emoji.Keycap1,
		"2":  emoji.Keycap2,
		"3":  emoji.Keycap3,
		"4":  emoji.Keycap4,
		"5":  emoji.Keycap5,
		"6":  emoji.Keycap6,
		"7":  emoji.Keycap7,
		"8":  emoji.Keycap8,
		"9":  emoji.Keycap9,
		"10": emoji.Keycap10,
	}

	displayPlayerSymbols := map[uint8]emoji.Emoji{
		EmptySymbol:   emoji.WhiteCircle,
		Player1Symbol: emoji.RedCircle,
		Player2Symbol: emoji.YellowCircle,
	}

	s := ""

	// print header
	for i := 0; i < len(b.Grid[0]); i++ {
		s += fmt.Sprintf("%v  ", displayTable[strconv.Itoa(i+1)])
	}
	s += "\n"

	// print board
	for _, row := range b.Grid {
		for _, col := range row {
			s += fmt.Sprintf("%v ", displayPlayerSymbols[col])
		}
		s += "\n"
	}

	return strings.TrimSpace(s)
}
