package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 3, 5, 7, 9, 11, 13, 15}

	for x, y := range a {
		fmt.Println(x, y)
	}

	var b = []int{2, 4, 6, 8, 10, 12, 14}

	for _, m := range b {
		println(m)
	}
}
