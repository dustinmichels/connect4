// Demo code for the Flex primitive.
package ui

import (
	"connect4/game"
	"fmt"

	"github.com/gdamore/tcell/v2"
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
	instructionsPanel.SetText(`* Use the arrow keys to move left and right`)

	// debug viz
	// boardWidget.SetBackgroundColor(tcell.ColorGrey)
	gameFrame := tview.NewFrame(boardWidget).
		SetBorders(1, 1, 1, 1, 2, 2)

	// gameFrame := tview.NewFrame(boardWidget).
	// 	SetBorders(2, 2, 2, 2, 4, 4).
	// 	AddText("Header [red]left[white]", true, tview.AlignLeft, tcell.ColorWhite).
	// 	AddText("Footer second middle", false, tview.AlignCenter, tcell.ColorGreen)

	// add borders and titles
	// playerWidget.SetBorder(true).SetTitle("Turn")
	// gameFrame.SetBorder(true).SetTitle("Game")
	// instructionsPanel.SetBorder(true).SetTitle("Instructions")
	// debugPanel.SetBorder(true).SetTitle("Debug")

	// blankPanel := tview.NewTextView().SetDynamicColors(true)

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	// menu := newPrimitive("Menu")
	// main := newPrimitive("Main content")
	// main := gameFrame

	leftSideBar := newPrimitive("")
	rightSideBar := newPrimitive("")

	header := tview.NewTextView().SetText(AsciiArt2).SetTextAlign(tview.AlignCenter)

	x, y, w, h := boardWidget.GetRect()
	fmt.Fprintf(debugPanel, "\n%v %v %v %v", x, y, w, h)

	grid := tview.NewGrid().
		SetRows(7, 0, 3).
		SetColumns(0, 25, 30, 0).
		SetBorders(true).
		AddItem(header, 0, 0, 1, 4, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 4, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.
		AddItem(leftSideBar, 1, 0, 1, 1, 0, 0, false).
		AddItem(gameFrame, 1, 1, 1, 1, 0, 0, false).
		AddItem(playerWidget, 1, 2, 1, 1, 0, 0, false).
		AddItem(rightSideBar, 1, 3, 1, 1, 0, 0, false)

	// Layout for screens wider than 100 cells.
	// grid.
	// 	AddItem(leftSideBar, 1, 0, 1, 1, 0, 100, false).
	// 	AddItem(gameFrame, 1, 1, 1, 1, 0, 100, false).
	// 	AddItem(playerWidget, 1, 2, 1, 1, 0, 100, false).
	// 	AddItem(rightSideBar, 1, 3, 1, 1, 0, 100, false)

	// Put together the layout
	flex := tview.NewFlex().
		AddItem(playerWidget, 20, 1, false).
		AddItem(gameFrame, 18, 1, false).
		AddItem(instructionsPanel, 40, 1, false).
		AddItem(debugPanel, 0, 1, false)

	flex.Box.SetBackgroundColor(tcell.NewHexColor(0x000000))

	// AddItem(debugPanel, 0, 1, false)

	// flex.SetBackgroundColor(tcell.NewHexColor(0x000000))

	if err := app.SetRoot(grid, true).EnableMouse(false).SetFocus(boardWidget).Run(); err != nil {
		panic(err)
	}

}
