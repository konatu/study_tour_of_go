package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []int{1, 3, 5, 7, 9}
	fmt.Println(s)

	s[1] = 5
	fmt.Println(s)
	printSlice(s)

	a := make([]int, 5, 10)
	printSlice(a)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	printSliceString(board)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceString(s [][]string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
