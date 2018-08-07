package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 5, 7, 9}
	fmt.Println(a)

	a[1] = 5
	fmt.Println(a)
}
