//adevent of code 2017, day 22, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	up = iota
	right
	down
	left
)

type vec2i struct {
	y, x int
}

var (
	grid1 = map[vec2i]rune{}
	grid2 = map[vec2i]rune{}
	moves = map[int]vec2i{
		up:    {-1, 0},
		right: {0, 1},
		down:  {1, 0},
		left:  {0, -1},
	}
)

func part1(y, x, dir int) (res int) {
	for i := 0; i < 10000; i++ {
		node := grid1[vec2i{y, x}]
		if node == '#' {
			dir = (dir + right) % len(moves)
			node = '.'
		} else {
			dir = (dir + left) % len(moves)
			node = '#'
			res++
		}
		grid1[vec2i{y, x}] = node
		y, x = y+moves[dir].y, x+moves[dir].x
	}
	return
}

func part2(y, x, dir int) (res int) {
	states := map[rune]rune{'.': 'W', 'W': '#', '#': 'F', 'F': '.'}
	for i := 0; i < 10000000; i++ {
		node := grid2[vec2i{y, x}]
		switch node {
		case 'W':
			node = states[node]
		case '#':
			dir = (dir + right) % len(moves)
			node = states[node]
		case 'F':
			dir = (dir + down) % len(moves)
			node = states[node]
		case '.':
			fallthrough
		default:
			dir = (dir + left) % len(moves)
			node = states['.'] // new node

		}
		if node == '#' {
			res++
		}
		grid2[vec2i{y, x}] = node
		y, x = y+moves[dir].y, x+moves[dir].x
	}
	return
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i, line := range lines {
		for j, c := range line {
			grid1[vec2i{i, j}] = c
			grid2[vec2i{i, j}] = c
		}
	}
	y, x, dir := len(lines)/2, len(lines[0])/2, up
	fmt.Println("part 1:", part1(y, x, dir))
	fmt.Println("part 2:", part2(y, x, dir))
}
