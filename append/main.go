package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)
	s = append(s, 0)

	s = append(s, 0, 1, 2, 3, 4)

	fmt.Println(s)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
