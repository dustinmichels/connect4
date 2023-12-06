// Demo code for the Flex primitive.
package ui

import (
	"connect4/game"

	"github.com/rivo/tview"
)

func initDebugPanel() *tview.TextView {
	debugPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).SetScrollable(true)
	return debugPanel
}

func Test() {

	// init game
	g := game.NewGame()
	b := g.Match.Board
	b.ApplyMoves([]int{0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 1, 3, 3})

	app := tview.NewApplication()

	debugPanel := initDebugPanel()
	playerWidget := initPlayerPanel()
	boardWidget := initBoardWidget(g, debugPanel, playerWidget)

	// rules panel
	instructionsPanel := tview.NewTextView().SetDynamicColors(true)
	instructionsPanel.SetText(`[green]Rules:[white]`)

	// debug viz
	// boardWidget.SetBackgroundColor(tcell.ColorGrey)
	gameFrame := tview.NewFrame(boardWidget).
		SetBorders(1, 1, 1, 1, 1, 1)

	// add borders and titles
	playerWidget.SetBorder(true).SetTitle("Turn")
	gameFrame.SetBorder(true).SetTitle("Game")
	instructionsPanel.SetBorder(true).SetTitle("Instructions")
	debugPanel.SetBorder(true).SetTitle("Debug")

	// Put together the layout
	flex := tview.NewFlex().
		AddItem(playerWidget, 20, 1, false).
		AddItem(gameFrame, 50, 1, false).
		AddItem(instructionsPanel, 50, 1, false).
		AddItem(debugPanel, 0, 1, false)

	// flex.SetBackgroundColor(tcell.NewHexColor(0x000000))

	if err := app.SetRoot(flex, true).EnableMouse(false).SetFocus(boardWidget).Run(); err != nil {
		panic(err)
	}

}
