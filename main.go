package main

import (
	"connect4/board"
	"fmt"
	"time"
)

const TestNum = 100_000

func TimeTestConcurrent() {
	board := board.NewBoard(7, 6)

	// time 100,000 loops
	start := time.Now()
	for i := 0; i < TestNum; i++ {
		board.GetWinnerConcurrent()
	}
	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Elapsed: %s\n", elapsed)
	fmt.Println("Average: ", elapsed/TestNum)
}

func TimeTestBasic() {
	board := board.NewBoard(7, 6)

	// time 100,000 loops
	start := time.Now()
	for i := 0; i < 100_000; i++ {
		board.GetWinner()
	}
	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Elapsed: %s\n", elapsed)
	fmt.Println("Average: ", elapsed/TestNum)
}

func BasicTest() {
	board := board.NewBoard(7, 6)
	board.ApplyMoves([]int{0, 1, 2, 3, 4, 5, 6, 0, 1, 2, 2, 3, 1, 3, 3})

	fmt.Println(board)

	winner, ok := board.GetWinnerConcurrent()
	if ok {
		fmt.Printf("Winner: %v\n", winner)
	} else {
		fmt.Println("No winner")
	}

}

func main() {

	fmt.Println("Basic:")
	TimeTestBasic()

	fmt.Println("\nWith concurrency:")
	TimeTestConcurrent()

	// BasicTest()
}
