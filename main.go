package main

import (
	"connect4/board"
	"fmt"
	"time"
)

func TimeTest() {

	board := board.NewBoard(7, 6)

	// time 100,000 loops
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		board.GetWinner()
	}
	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Elapsed: %s\n", elapsed)

}

func main() {

	TimeTest()

	// board := board.NewBoard(7, 6)

	// // board.ApplyMoves([]int{0, 1, 2, 3, 4, 5, 6, 0, 1, 2})
	// board.ApplyMoves([]int{0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 2, 3, 1, 3, 3})

	// fmt.Println(board)

	// winner, ok := board.GetWinner()
	// if ok {
	// 	fmt.Printf("Winner: %v\n", winner)
	// } else {
	// 	fmt.Println("No winner")
	// }

}
