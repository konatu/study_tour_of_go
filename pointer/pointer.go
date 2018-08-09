// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	i, j := 10, 20
//
// 	f := &i
// 	fmt.Println(f)
// 	fmt.Println(i)
// 	fmt.Println(&i)
// 	fmt.Println(*f)
//
// 	p := &j
// 	fmt.Println(j)
// 	fmt.Println(&j)
// 	fmt.Println(p)
// 	fmt.Println(*p)
// 	*p = 5
// 	fmt.Println(p)
// 	fmt.Println(*p)
// }

package main

import (
	"fmt"
)

var a string = "test"

func main() {
	fmt.Println(a)
	fmt.Println(&a) //a *sring
	s := &a         //a *sring
	fmt.Println(s)
	fmt.Println(*s)
}
