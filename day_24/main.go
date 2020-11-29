//advent of code 2017, day 24, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	left, right int
	inUse       bool
}

var nodes = map[*node]bool{}
var part1, part2, part2L int

func search(port, strength, length int) {
	if strength > part1 {
		part1 = strength
	}
	if strength > part2 {
		if length >= part2L {
			part2L = length
			part2 = strength
		}
	}
	for node := range nodes {
		if node.inUse {
			continue
		}
		if !(node.left == port || node.right == port) {
			continue
		}
		node.inUse = true
		s := node.left + node.right
		if port == node.left {
			search(node.right, strength+s, length+1)
		} else {
			search(node.left, strength+s, length+1)
		}
		node.inUse = false
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		split := strings.Split(line, "/")
		left, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		nodes[&node{left, right, false}] = true
	}
	search(0, 0, 0)
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}
