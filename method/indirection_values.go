package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Ads() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Adsfunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Ads())
	fmt.Println(Adsfunc(v))

	p := Vertex{4, 3}
	fmt.Println(p.Ads())
	fmt.Println(Adsfunc(p))
}
