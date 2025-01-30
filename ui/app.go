package ui

import (
	"connect4/game"
	"fmt"

	"github.com/rivo/tview"
)

const debug = false

func initDebugPanel() *tview.TextView {
	debugPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).SetScrollable(true)
	return debugPanel
}

func StartApp(g *game.Game) {

	app := tview.NewApplication()

	pages := tview.NewPages()
	debugPanel := initDebugPanel()
	playerWidget := initPlayerPanel()
	boardWidget := initBoardWidget(g, debugPanel, playerWidget, pages)

	gameFrame := tview.NewFrame(boardWidget).
		SetBorders(1, 1, 1, 1, 2, 2)

	leftSideBar := tview.NewBox()
	rightSideBar := tview.NewBox()

	footerText := `arrow keys + <space> or <enter> to select a column | ctrl-c to quit`

	header := tview.NewTextView().SetText(AsciiArt2).SetTextAlign(tview.AlignCenter)
	footer := tview.NewTextView().SetText(footerText).SetTextAlign(tview.AlignCenter)

	footer.SetBorder(true).SetTitle(" Guide ")

	// print dimensions to debug panel
	x, y, w, h := boardWidget.GetRect()
	fmt.Fprintf(debugPanel, "\n%v %v %v %v", x, y, w, h)

	grid := tview.NewGrid().
		SetRows(8, 0, 10, 0, 3).
		SetColumns(0, 25, 25, 0).
		// SetBorders(true).
		AddItem(header, 0, 0, 1, 4, 0, 0, false).
		AddItem(footer, 4, 0, 1, 4, 0, 0, false)

	grid.
		AddItem(tview.NewBox(), 1, 0, 1, 4, 0, 0, false). // upper fill
		AddItem(leftSideBar, 2, 0, 1, 1, 0, 0, false).
		AddItem(gameFrame, 2, 1, 1, 1, 0, 0, true).
		AddItem(playerWidget, 2, 2, 1, 1, 0, 0, false).
		AddItem(rightSideBar, 2, 3, 1, 1, 0, 0, false).
		AddItem(tview.NewBox(), 3, 0, 1, 4, 0, 0, false)

	if debug {
		grid.AddItem(debugPanel, 1, 0, 1, 1, 0, 0, false) // lower fill
	}

	modal := tview.NewModal().
		SetText("Game over!").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})

	// ----- MENU PAGE -----
	startSinglePlayer := func() {
		pages.SwitchToPage("grid")
		// app.SetFocus(boardWidget) // Set focus on BoardWidget when transitioning
	}
	menu := makeMenuPage(app, startSinglePlayer)

	pages.AddPage("menu", menu, true, true).
		AddPage("grid", grid, true, false).
		AddPage("modal", modal, false, false)

	if err := app.SetRoot(pages, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}

}
