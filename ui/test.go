// Demo code for the Flex primitive.
package ui

import (
	"connect4/game"
	"fmt"

	"github.com/rivo/tview"
)

func initDebugPanel() *tview.TextView {
	debugPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).SetScrollable(true)
	return debugPanel
}

func Test() {

	// modalShown := false
	pages := tview.NewPages()

	// init game
	g := game.NewGame()
	// b := g.Match.Board
	// b.ApplyMoves([]int{ 0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 1, 3, 3})

	app := tview.NewApplication()

	debugPanel := initDebugPanel()
	playerWidget := initPlayerPanel()
	boardWidget := initBoardWidget(g, debugPanel, playerWidget, pages)

	// rules panel
	instructionsPanel := tview.NewTextView().SetDynamicColors(true)
	instructionsPanel.SetText(`* Use the arrow keys to move left and right`)

	// debug viz
	// boardWidget.SetBackgroundColor(tcell.ColorGrey)
	gameFrame := tview.NewFrame(boardWidget).
		SetBorders(1, 1, 1, 1, 2, 2)

	// // add borders and titles
	// playerWidget.SetBorder(true).SetTitle("Turn")
	// gameFrame.SetBorder(true).SetTitle("Game")
	// instructionsPanel.SetBorder(true).SetTitle("Instructions")
	// debugPanel.SetBorder(true).SetTitle("Debug")

	leftSideBar := tview.NewBox()
	rightSideBar := tview.NewBox()

	footerText := `use arrow keys to navigate | <space> or <enter> to select a column | ctrl-c to quit`

	header := tview.NewTextView().SetText(AsciiArt2).SetTextAlign(tview.AlignCenter)
	footer := tview.NewTextView().SetText(footerText).SetTextAlign(tview.AlignCenter)

	footer.SetBorder(true).SetTitle(" Guide ")

	x, y, w, h := boardWidget.GetRect()
	fmt.Fprintf(debugPanel, "\n%v %v %v %v", x, y, w, h)

	grid := tview.NewGrid().
		SetRows(7, 0, 10, 0, 3).
		SetColumns(0, 25, 25, 0).
		SetBorders(false).
		AddItem(header, 0, 0, 1, 4, 0, 0, false).
		AddItem(footer, 4, 0, 1, 4, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.
		AddItem(tview.NewBox(), 1, 0, 1, 4, 0, 0, false). // upper fill
		AddItem(leftSideBar, 2, 0, 1, 1, 0, 0, false).
		AddItem(gameFrame, 2, 1, 1, 1, 0, 0, false).
		AddItem(playerWidget, 2, 2, 1, 1, 0, 0, false).
		AddItem(rightSideBar, 2, 3, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 3, 0, 1, 4, 0, 0, false) // lower fill

	// Layout for screens wider than 100 cells.
	// grid.
	// 	AddItem(tview.NewBox(), 1, 0, 1, 4, 0, 100, false). // upper fill
	// 	AddItem(leftSideBar, 2, 0, 1, 1, 0, 100, false).
	// 	AddItem(gameFrame, 2, 1, 1, 1, 0, 100, false).
	// 	AddItem(playerWidget, 2, 2, 1, 1, 0, 100, false).
	// 	AddItem(rightSideBar, 2, 3, 1, 1, 0, 100, false).
	// 	AddItem(tview.NewBox(), 3, 0, 1, 4, 0, 100, false) // lower fill

	// Put together the layout
	// flex := tview.NewFlex().
	// 	AddItem(playerWidget, 20, 1, false).
	// 	AddItem(gameFrame, 18, 1, false).
	// 	AddItem(instructionsPanel, 40, 1, false).
	// 	AddItem(debugPanel, 0, 1, false)

	// flex.Box.SetBackgroundColor(tcell.NewHexColor(0x000000))

	// AddItem(debugPanel, 0, 1, false)

	// flex.SetBackgroundColor(tcell.NewHexColor(0x000000))

	modal := tview.NewModal().
		SetText("Game over!").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})

	// modal.

	pages.AddPage("grid", grid, true, true).
		AddPage("modal", modal, false, false)

	// pages.ShowPage("modal")
	// pages.SwitchToPage("modal")

	if err := app.SetRoot(pages, true).EnableMouse(false).SetFocus(boardWidget).Run(); err != nil {
		panic(err)
	}

}
