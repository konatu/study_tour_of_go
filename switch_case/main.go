package main

import (
	"fmt"
)

func main() {
	i := 2
	switch i {
	case 0:
		fmt.Println("out", i)
	case 1:
		fmt.Println("out", i)
	case 2:
		fmt.Println("ok", i)
	}
}
