package main

import (
	"connect4/bot"
	"connect4/game"
	"fmt"
)

func main() {
	g := game.NewGame()
	b := g.Match.Board
	b.ApplyMoves([]int{0, 1, 0, 2, 0, 2, 4})
	fmt.Println(b)
	fmt.Println(bot.Score(b))

	g = game.NewGame()
	b = g.Match.Board
	b.ApplyMoves([]int{0, 1, 0, 2, 0, 2, 1})
	fmt.Println(b)
	fmt.Println(bot.Score(b))

}
