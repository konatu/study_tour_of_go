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

var t map[string]trim

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["RedBull"] = Vertex{
		40.55554, 45.68443,
	}
	fmt.Println(m["RedBull"])

	t = make(map[string]trim)
	t["Monster"] = trim{
		"enagy", "drink",
	}
	fmt.Println(t["Monster"])
}
