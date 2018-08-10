package main

import (
	"fmt"
	"sync"
	"time"
)

type Safecounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Safecounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *Safecounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := Safecounter{v: make(map[string]int)}
	for i := 0; i < 550; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
