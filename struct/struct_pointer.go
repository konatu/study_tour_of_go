package main

import (
	"fmt"
)

type st struct {
	X int
	Y int
}

func main() {
	v := st{10, 20}
	p := &v
	p.Y = 30 // (*p).Y
	fmt.Println(*p)
}
