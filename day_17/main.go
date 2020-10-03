//advent of code 2017, day 17, part 1
package main

import (
	"container/ring"
	"fmt"
)

const puzzle = 337

func main() {
	r := ring.New(1)
	r.Value = 0
	for i := 1; i < 2018; i++ {
		r = r.Move(puzzle)
		s := ring.New(1)
		s.Value = i
		r = s.Link(r).Next()
	}
	fmt.Printf("part 1: %d\n", r.Value.(int))
}
