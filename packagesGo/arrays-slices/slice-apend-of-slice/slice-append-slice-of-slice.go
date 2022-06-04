package main

import (
	"fmt"
	"strings"
)

func main() {
	//Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The playes take runs.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	fmt.Println(len(board), cap(board))

	for i := 0; i < len(board); i++ {
		// fmt.Println(board[i])
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	t := make([]int, len(s), (cap(s)+1)*2)
	copy(t, s)
	printSlice(t)

	s = append(s, t...)
	printSlice(s)

	t = append(t, s...)
	printSlice(t)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
