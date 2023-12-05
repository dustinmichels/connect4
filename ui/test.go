// Demo code for the Flex primitive.
package main

import (
	"connect4/game"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func drawBoard(b *game.Board, table *tview.Table) {

	symbolMap := map[string]string{
		game.EmptySymbol:   "●",
		game.Player1Symbol: "[red]x[white]",
		game.Player2Symbol: "[darkblue]o[white]",
		"fancy":            "●",
	}

	// add header row
	for i := 0; i < b.NumCols(); i++ {
		table.SetCell(0, i,
			tview.NewTableCell(fmt.Sprintf("%d", i+1)).
				SetAlign(tview.AlignCenter))
		table.SetCell(1, i,
			tview.NewTableCell("-").
				SetAlign(tview.AlignCenter))
	}

	// set up cells
	for i := 0; i < b.NumRows(); i++ {
		for j := 0; j < b.NumCols(); j++ {
			symbol := symbolMap[b.Get(i, j)]
			table.SetCell(i+2, j,
				tview.NewTableCell(symbol).
					// SetBackgroundColor(tcell.ColorGrey).
					// SetTextColor(tcell.ColorRed).
					SetAlign(tview.AlignCenter))
		}
	}

}

func initBoardPanel(b *game.Board) *tview.Table {
	table := tview.NewTable().SetBorders(false).SetSelectable(true, true)
	drawBoard(b, table)
	return table
}

func main() {
	app := tview.NewApplication()

	// set up player panel
	playerPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetTextAlign(tview.AlignCenter)

	fmt.Fprintf(playerPanel, `%s["one"]Player 1[""]%s["two"]Player 2[""]`, "\n", "\n")
	playerPanel.Highlight("one")

	// debug
	debugPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).SetScrollable(true)

	// set up board
	b := game.NewBoard()
	b.ApplyMoves([]int{0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 2, 3, 1, 3, 3})
	boardPanel := initBoardPanel(b)

	// debug viz
	boardPanel.SetBackgroundColor(tcell.ColorGrey)
	gameFrame := tview.NewFrame(boardPanel).
		SetBorders(1, 1, 1, 1, 1, 1)

	// capture input
	boardPanel.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		boardPanel.SetSelectable(true, true)
		_, col := boardPanel.GetSelection()
		switch event.Key() {
		case tcell.KeyLeft:
			fmt.Fprintf(debugPanel, "\nLeft pressed. %v -> %v", col, col-1)
			if col > 0 {
				boardPanel.Select(0, col-1)
			}
		case tcell.KeyRight:
			fmt.Fprintf(debugPanel, "\nRight pressed. %v -> %v", col, col+1)
			if col < b.NumCols()-1 {
				boardPanel.Select(0, col+1)
			}
		case tcell.KeyEnter:
			return event
		}
		return nil
	})

	boardPanel.SetSelectedFunc(func(row int, col int) {
		fmt.Fprintf(debugPanel, "\nSelected %v, %v", row, col)
		boardPanel.SetSelectable(false, true)
		err := b.Update(true, col)
		if err != nil {
			fmt.Fprintf(debugPanel, "\n%v", err)
		}
		drawBoard(b, boardPanel)
	})

	// add borders and titles
	playerPanel.SetBorder(true).SetTitle("Turn")
	// boardPanel.SetBorder(true).SetTitle("Board")
	gameFrame.SetBorder(true).SetTitle("Game")
	debugPanel.SetBorder(true).SetTitle("Debug")

	// Put together the layout
	flex := tview.NewFlex().
		AddItem(playerPanel, 20, 1, false).
		AddItem(gameFrame, 50, 1, false).
		AddItem(debugPanel, 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(false).SetFocus(boardPanel).Run(); err != nil {
		panic(err)
	}

}
