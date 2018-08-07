package main

import (
	"fmt"
)

func main() {
	a := [6]int{1, 3, 5, 7, 9, 11}
	fmt.Println(a)

	var s []int = a[1:4]
	fmt.Println(s)
}
