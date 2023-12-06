package main

import (
	"connect4/bot"
	"connect4/game"
	"fmt"
)

func main() {
	g := game.NewGame()
	b := g.Match.Board

	b.ApplyMoves([]int{0, 1, 0, 2, 0, 2, 0})

	fmt.Println(b)

	score := bot.Score(b)

	fmt.Println(score)

}
