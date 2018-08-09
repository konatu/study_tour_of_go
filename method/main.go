package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type myinit int

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f myinit) Bbs() int {
	return int(f)
}

func (v *Vertex) scale(t float64) {
	v.X = v.X * t
	v.Y = v.Y * t
}

func main() {
	v := Vertex{3, 4}
	v.scale(10)
	fmt.Println(v.Abs())

	f := myinit(7)
	fmt.Println(f.Bbs())
}

// methodを使わない場合
//
// func abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
//
// func main() {
// 	v := Vertex{3, 4}
// 	fmt.Println(abs(v))
// }
