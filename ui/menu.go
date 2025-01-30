package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// makeMenuPage creates a menu page with arrow key navigation
func makeMenuPage(app *tview.Application, buttonAction func()) tview.Primitive {
	menu := tview.NewFlex().SetDirection(tview.FlexRow)
	menu.SetTitle("Menu")

	header := tview.NewTextView().SetText(AsciiArt2).SetTextAlign(tview.AlignCenter)

	// Create buttons
	buttons := []*tview.Button{
		tview.NewButton("Singleplayer (vs Bot)").SetSelectedFunc(buttonAction),
		tview.NewButton("Multiplayer (local)").SetSelectedFunc(buttonAction),
		tview.NewButton("Multiplayer (online)").SetSelectedFunc(buttonAction),
		tview.NewButton("Run server").SetSelectedFunc(buttonAction),
	}

	menu.AddItem(header, 0, 1, false)
	for _, btn := range buttons {
		menu.AddItem(btn, 1, 1, true)
	}
	menu.AddItem(tview.NewBox(), 0, 1, false)

	// Enable arrow key navigation
	currentIndex := 0
	menu.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			currentIndex = (currentIndex - 1 + len(buttons)) % len(buttons)
			app.SetFocus(buttons[currentIndex])
			return nil
		case tcell.KeyDown:
			currentIndex = (currentIndex + 1) % len(buttons)
			app.SetFocus(buttons[currentIndex])
			return nil
		}
		return event
	})

	return menu
}
