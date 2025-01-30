package ui

import (
	"github.com/rivo/tview"
)

// makeMenuPage creates a menu page with arrow key navigation
func makeMenuPage(app *tview.Application, buttonAction func()) tview.Primitive {
	menu := tview.NewFlex().SetDirection(tview.FlexRow)
	menu.SetTitle("Menu")

	header := tview.NewTextView().SetText(AsciiArt2).SetTextAlign(tview.AlignCenter)

	list := tview.NewList().
		AddItem("Play local", "1-v-1 and 1-v-Bot", '1', buttonAction).
		AddItem("Play online", "Host or join a server", '2', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	list.SetHighlightFullLine(true)
	list.SetBorder(true).SetTitle(" Welcome ")

	listCentered := tview.NewFlex().SetDirection(tview.FlexColumn)
	listCentered.AddItem(tview.NewBox(), 0, 1, false)
	listCentered.AddItem(list, 0, 3, true)
	listCentered.AddItem(tview.NewBox(), 0, 1, false)

	menu.AddItem(header, 0, 1, false)
	menu.AddItem(listCentered, 0, 1, true)
	menu.AddItem(tview.NewBox(), 0, 1, false)

	return menu
}
