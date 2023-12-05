// Demo code for the Flex primitive.
package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func initBoard() *tview.Table {
	// set up board
	board := tview.NewTable().SetBorders(true)

	board.SetCell(0, 0,
		tview.NewTableCell("a").
			SetTextColor(tcell.ColorRed).
			SetAlign(tview.AlignCenter))

	board.SetCell(0, 1,
		tview.NewTableCell("a").
			SetTextColor(tcell.ColorRed).
			SetAlign(tview.AlignCenter))

	return board
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

	// set up board
	boardPanel := initBoard()

	// add borders and titles
	playerPanel.SetBorder(true).SetTitle("Player")
	boardPanel.SetBorder(true).SetTitle("Game")

	// Put together the layout
	flex := tview.NewFlex().
		AddItem(playerPanel, 20, 1, false).
		AddItem(boardPanel, 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
