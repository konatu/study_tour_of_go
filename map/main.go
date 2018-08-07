package main

import (
	"fmt"
)

type Vertex struct {
	lat, long float64
}

type trim struct {
	x, y string
}

var m = map[string]Vertex{
	"RedBull": Vertex{
		40.55554, 45.68443,
	}, "Monster": Vertex{
		45.88634, 58.4479,
	},
}

func main() {
	fmt.Println(m)
}
