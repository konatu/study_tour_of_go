package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 5
	ch <- 15
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
