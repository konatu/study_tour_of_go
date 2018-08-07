package main

import (
	"fmt"
	"time"
)

func main() {
	i := time.Now()
	switch {
	case i.Hour() < 12:
		fmt.Println("good mornig", i)
	case i.Hour() < 17:
		fmt.Println("hello", i)
	default:
		fmt.Println("good evening", i)
	}
}
