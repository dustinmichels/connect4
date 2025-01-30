package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	// Create a menu page
	menuPage := makeMenuPage(app, func() {
		app.Stop()
	})

	// Set the root page
	app.SetRoot(menuPage, true)

	// Start the app
	if err := app.Run(); err != nil {
		panic(err)
	}
}

// makeMenuPage creates a menu page with arrow key navigation
func makeMenuPage(app *tview.Application, buttonAction func()) tview.Primitive {
	menu := tview.NewFlex().SetDirection(tview.FlexRow)

	// Create buttons
	buttons := []*tview.Button{
		tview.NewButton("Btn 1").SetSelectedFunc(buttonAction),
		tview.NewButton("Btn 2").SetSelectedFunc(buttonAction),
		tview.NewButton("Btn 3").SetSelectedFunc(buttonAction),
		tview.NewButton("Btn 4").SetSelectedFunc(buttonAction),
	}

	// Apply colors to buttons
	for _, btn := range buttons {
		btn.SetLabelColor(tcell.ColorWhite.TrueColor())                // Text color
		btn.SetBackgroundColor(tcell.ColorBlack.TrueColor())           // Default background color
		btn.SetLabelColorActivated(tcell.ColorRed.TrueColor())         // Text color when selected
		btn.SetBackgroundColorActivated(tcell.ColorYellow.TrueColor()) // Background when selected
	}

	// Add buttons to menu
	for _, btn := range buttons {
		menu.AddItem(btn, 1, 1, true)
	}

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
