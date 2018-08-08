package main

import (
	"fmt"
)

func decorate(a string) string {
	return "<<" + a + ">>"
}

func main() {
	f := decorate
	ret := f("golang")
	fmt.Println(ret)
}

// func main() {
// 	fmt.Println(decorate("golang"))
// }
