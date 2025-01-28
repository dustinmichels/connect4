// Demo code for the Flex primitive.
package ui

import (
	"connect4/game"
)

func Test() {

	// init game
	g := game.NewGame()

	// b := g.Match.Board
	// b.ApplyMoves([]int{0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 3, 1, 3, 3})

	initApp(g)
}
