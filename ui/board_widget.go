package ui

import (
	"connect4/game"
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func initBoardWidget(g *game.Game, debugPanel *tview.TextView, playerWidget *tview.TextView) *tview.Table {
	b := g.Match.Board

	boardWidget := tview.NewTable().SetBorders(false).SetSelectable(true, true)
	updateBoardWidget(b, boardWidget)

	// boardWidget.SetBorders(true)

	artificialEnter := tcell.NewEventKey(tcell.KeyEnter, 0, tcell.ModNone)

	// Capture keystrokes
	boardWidget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		boardWidget.SetSelectable(true, true)
		_, col := boardWidget.GetSelection()

		switch event.Key() {
		case tcell.KeyLeft:
			fmt.Fprintf(debugPanel, "\nLeft pressed. %v -> %v", col, col-1)
			if col > 0 {
				boardWidget.Select(0, col-1)
			}
		case tcell.KeyRight:
			fmt.Fprintf(debugPanel, "\nRight pressed. %v -> %v", col, col+1)
			if col < b.NumCols()-1 {
				boardWidget.Select(0, col+1)
			}
		case tcell.KeyEnter:
			return event
		default:
			fmt.Fprintf(debugPanel, "\nKey pressed: %v", event.Name())
		}

		// check for number key. if found, also submit an "enter" event
		rune := event.Rune()
		runeInt, err := strconv.Atoi(string(rune))
		if err == nil && runeInt > 0 && runeInt <= b.NumCols() {
			fmt.Fprintf(debugPanel, "\n%v pressed. %v -> %v", rune, col, runeInt-1)
			boardWidget.Select(0, runeInt-1)
			return artificialEnter
		}

		// space bar is handled like enter
		if event.Key() == tcell.KeyRune && rune == ' ' {
			return artificialEnter
		}

		return nil
	})

	// Handle selection (enter pressed while a cell is selected)
	boardWidget.SetSelectedFunc(func(row int, col int) {
		fmt.Fprintf(debugPanel, "\nSelected %v, %v", row, col)
		// boardPanel.SetSelectable(false, true)

		err := g.Match.ApplyMove(col)
		if err != nil {
			fmt.Fprintf(debugPanel, "\n%v", err)
		}

		updatePlayerPanel(playerWidget, g.Match.IsPlayer1Active())
		updateBoardWidget(b, boardWidget)

		winner, found := g.Match.Board.GetWinner()
		if found {
			fmt.Fprintf(debugPanel, "\nPlayer %v wins!", winner)
		}

	})

	return boardWidget
}

func updateBoardWidget(b *game.Board, boardWidget *tview.Table) {

	player1Color, _ := MakeTcellColor(AnsiRed)
	player2Color := tcell.NewHexColor(AnsiYellowHex)
	defaultBgColor := tcell.NewHexColor(AnsiBlueHex)

	symbolMap := map[string]string{
		game.EmptySymbol:   "●",
		game.Player1Symbol: fmt.Sprintf("[%s]X[white]", "lightgrey"),
		game.Player2Symbol: fmt.Sprintf("[%s]0[white]", "gray"),
		"fancy":            "●",
	}

	// add header row
	for i := 0; i < b.NumCols(); i++ {
		boardWidget.SetCell(0, i,
			tview.NewTableCell(fmt.Sprintf(" %d", i+1)).
				SetAlign(tview.AlignCenter))
		boardWidget.SetCell(1, i,
			tview.NewTableCell("-").
				SetAlign(tview.AlignCenter))
	}

	// set up cells
	for i := 0; i < b.NumRows(); i++ {
		for j := 0; j < b.NumCols(); j++ {
			symbol := b.Get(i, j)
			drawSymbol := symbolMap[symbol]
			bgColor := defaultBgColor
			if symbol == "X" {
				bgColor = player1Color
			}
			if symbol == "O" {
				bgColor = player2Color
			}

			boardWidget.SetCell(i+2, j,
				tview.NewTableCell(drawSymbol).
					SetBackgroundColor(bgColor).
					// SetTextColor(tcell.ColorRed).
					SetAlign(tview.AlignCenter))
		}
	}

}
