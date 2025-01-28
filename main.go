package main

import (
	"connect4/game"
	"connect4/ui"
)

func main() {
	g := game.NewGame()
	ui.StartApp(g)
}
