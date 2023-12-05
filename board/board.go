package board

import (
	"fmt"
	"strconv"

	"github.com/enescakir/emoji"
)

const EmptySymbol = "-"
const Player1Symbol = "X"
const Player2Symbol = "O"

type Board struct {
	grid [][]string
}

func NewBoard() *Board {
	width := 7
	height := 6
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = EmptySymbol
		}
	}
	return &Board{grid}
}

func NewCustomBoard(width, height int) *Board {

	if width > 10 {
		panic("Width must be less than 10")
	}

	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
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

func (b *Board) ApplyMoves(moves []int) error {
	for i, col := range moves {
		if err := b.Update(i%2 == 0, col); err != nil {
			return err
		}
	}
	return nil
}

func (b *Board) String() string {

	displayTable := map[string]emoji.Emoji{
		EmptySymbol:   emoji.WhiteCircle,
		Player1Symbol: emoji.RedCircle,
		Player2Symbol: emoji.BlueCircle,
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
