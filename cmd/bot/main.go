package main

import (
	"connect4/game"
)

func main() {
	g := game.NewGame()
	b := g.Match.Board
	b.Print()
}
