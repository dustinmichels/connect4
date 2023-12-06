package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

func initPlayerPanel() *tview.TextView {

	// set up player panel
	playerPanel := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetTextAlign(tview.AlignCenter)

	// fmt.Fprintf(playerPanel, `%s["one"]Player 1 [red]x[white][""]%s["two"]Player 2 [yellow]o[white][""]`, "\n", "\n")
	// playerPanel.Highlight("one")

	updatePlayerPanel(playerPanel, true)

	return playerPanel

}

func updatePlayerPanel(playerPanel *tview.TextView, player1Active bool) {

	playerPanel.Clear()

	player1Label := fmt.Sprintf(`["one"]Player 1 [%s]x[white][""]`, Player1Color)
	player2Label := fmt.Sprintf(`["two"]Player 2 [%s]o[white][""]`, Player2Color)

	// set up player panel
	fmt.Fprintf(playerPanel, "\n%s\n%s", player1Label, player2Label)
	if player1Active {
		playerPanel.Highlight("one")
	} else {
		playerPanel.Highlight("two")
	}

}
