package main

import (
	"fmt"
)

func main() {
	enagy := [10]string{"monster", "redbull", "panther", "star"}
	fmt.Println(enagy)

	drink := enagy[1:3]
	fmt.Println(drink)

	strong := enagy[0:2]
	fmt.Println(strong)

	strong[1] = "coca"
	fmt.Println(enagy)
}
