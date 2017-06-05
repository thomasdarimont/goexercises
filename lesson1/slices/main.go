package main

import (
	"fmt"
	"strings"
)

func main() {
	simpleSliceDemo()

	nestedSliceDemo()

	appendDemo()

}
func appendDemo() {

	var s []int
	fmt.Println(s)

	// Add one item to the slice
	s = append(s, 1)
	fmt.Println(s)

	// Add more than on item to the slice
	s = append(s, 2, 3, 4)
	fmt.Println(s)
}

func nestedSliceDemo() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func simpleSliceDemo() {
	// The make function allocates a zeroed array and returns a slice that refers to that array
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
