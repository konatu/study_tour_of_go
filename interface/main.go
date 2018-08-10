// package main
//
// import (
// 	"fmt"
// 	"math"
// )
//
// type abser interface {
// 	Abs() float64
// }
//
// type MyFloat float64
//
// type Vertex struct {
// 	X, Y float64
// }
//
// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }
//
// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }
//
// func main() {
// 	var a abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}
//
// 	a = f
// 	a = &v
//
// 	fmt.Println(a.Abs())
// }

package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
