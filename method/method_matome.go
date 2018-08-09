package main

import (
	"fmt"
)

type calc struct {
	value1 int
	value2 int
}

type sum []calc

func (a calc) sum() int {
	return a.value1 + a.value2
}

func (s sum) sum() int {
	answer := 0
	for _, s := range s {
		answer += s.sum()
	}
	return answer
}

func (b calc) multi() int {
	return b.value1 * b.value2
}

func main() {
	sums := sum{
		{8, 5},
		{4, 5},
		{8, 6},
	}
	fmt.Println(sums.sum())

	s := calc{10, 20}
	fmt.Println(s.sum())
	fmt.Println(s.multi())
}
