package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := map[int]bool{}
	rand.Seed(time.Now().UnixNano())

	for len(r) < 6 {
		r[rand.Intn(45)+1] = true
	}

	for x := range r {
		fmt.Println(x)
	}
}
