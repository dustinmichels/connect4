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

	updatePlayerPanel(playerPanel, true)
	return playerPanel
}

func updatePlayerPanel(playerPanel *tview.TextView, player1Active bool) {

	playerPanel.Clear()

	startSymbol1 := ""
	endSymbol1 := ""
	startSymbol2 := ""
	endSymbol2 := ""

	if player1Active {
		startSymbol1 = "->"
		endSymbol1 = "  "
	} else {
		startSymbol2 = "->"
		endSymbol2 = "  "
	}

	player1Label := fmt.Sprintf(`["one"]%s Player 1 [%s]x[white] %s[""]`, startSymbol1, Player1Color, endSymbol1)
	player2Label := fmt.Sprintf(`["two"]%s Player 2 [%s]o[white] %s[""]`, startSymbol2, Player2Color, endSymbol2)

	// set up player panel
	fmt.Fprintf(playerPanel, "\n%s\n%s", player1Label, player2Label)

	// if player1Active {
	// 	playerPanel.Highlight("one")
	// } else {
	// 	playerPanel.Highlight("two")
	// }

}
