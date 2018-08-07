package main

import (
	"fmt"
)

func main() {
	var a [10]string
	a[0] = "red"
	a[1] = "bull"

	fmt.Println(a[0], a[1])

	enagy := [6]string{"redbull", "monster", "panther"}
	fmt.Println(enagy)
}
