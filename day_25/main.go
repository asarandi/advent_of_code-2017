//advent of code 2017, day 25, part 1
package main

import "fmt"

const left, right = -1, 1

type action struct {
	value, move int
	state       rune
}

var states = map[rune][]action{
	'A': {{1, right, 'B'}, {1, left, 'E'}},
	'B': {{1, right, 'C'}, {1, right, 'F'}},
	'C': {{1, left, 'D'}, {0, right, 'B'}},
	'D': {{1, right, 'E'}, {0, left, 'C'}},
	'E': {{1, left, 'A'}, {0, right, 'D'}},
	'F': {{1, right, 'A'}, {1, right, 'C'}},
}

func main() {
	const steps = 12459852
	state, cursor, res := 'A', 0, 0
	tape := make(map[int]int)
	for i := 0; i < steps; i++ {
		s := states[state][tape[cursor]]
		tape[cursor] = s.value
		cursor += s.move
		state = s.state
	}
	for _, value := range tape {
		res += value
	}
	fmt.Println("part 1:", res)
}
