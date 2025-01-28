package game

// CheckWinner checks if there is a winner on the board.
// Returns the symbol of the winner (Player1Symbol or Player2Symbol) or EmptySymbol if no winner.
func checkWinner(b *Board) uint8 {
	// Check horizontal rows
	for row := 0; row < boardHeight; row++ {
		for col := 0; col <= boardWidth-winningLength; col++ {
			if winner := b.checkLine(b.Grid[row][col : col+winningLength]); winner != EmptySymbol {
				return winner
			}
		}
	}

	// Check vertical columns
	for col := 0; col < boardWidth; col++ {
		for row := 0; row <= boardHeight-winningLength; row++ {
			if winner := b.checkVertical(row, col); winner != EmptySymbol {
				return winner
			}
		}
	}

	// Check diagonals (top-left to bottom-right)
	for row := 0; row <= boardHeight-winningLength; row++ {
		for col := 0; col <= boardWidth-winningLength; col++ {
			if winner := b.checkDiagonal(row, col, 1, 1); winner != EmptySymbol {
				return winner
			}
		}
	}

	// Check diagonals (bottom-left to top-right)
	for row := winningLength - 1; row < boardHeight; row++ {
		for col := 0; col <= boardWidth-winningLength; col++ {
			if winner := b.checkDiagonal(row, col, -1, 1); winner != EmptySymbol {
				return winner
			}
		}
	}

	return EmptySymbol // No winner found
}

// checkLine checks if all elements in a slice are the same and not EmptySymbol.
func (b *Board) checkLine(line []uint8) uint8 {
	symbol := line[0]
	if symbol == EmptySymbol {
		return EmptySymbol
	}
	for _, cell := range line {
		if cell != symbol {
			return EmptySymbol
		}
	}
	return symbol
}

// checkVertical checks for a vertical sequence of four matching symbols starting at (row, col).
func (b *Board) checkVertical(row, col int) uint8 {
	symbol := b.Grid[row][col]
	if symbol == EmptySymbol {
		return EmptySymbol
	}
	for i := 1; i < winningLength; i++ {
		if b.Grid[row+i][col] != symbol {
			return EmptySymbol
		}
	}
	return symbol
}

// checkDiagonal checks for a diagonal sequence of four matching symbols starting at (row, col).
// The parameters rowStep and colStep determine the direction of the diagonal.
func (b *Board) checkDiagonal(row, col, rowStep, colStep int) uint8 {
	symbol := b.Grid[row][col]
	if symbol == EmptySymbol {
		return EmptySymbol
	}
	for i := 1; i < winningLength; i++ {
		if b.Grid[row+i*rowStep][col+i*colStep] != symbol {
			return EmptySymbol
		}
	}
	return symbol
}
