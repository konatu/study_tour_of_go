package main

import (
	"fmt"
)

func main() {
	i, j := 10, 20

	f := &i
	fmt.Println(f)

	p := &j
	fmt.Println(*p)
}
