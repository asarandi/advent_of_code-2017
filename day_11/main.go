// advent of code 2017: day 11, part 1 and 2

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func dist(y, x int) int {
	y, x = abs(y), abs(x)
	if y < x {
		y, x = x, y
	}
	return (y-x)/2 + x
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	steps := strings.Split(strings.TrimSpace(string(data)), ",")
	y, x, max := 0, 0, 0
	for _, s := range steps {
		switch s {
		case "n":
			y -= 2
		case "s":
			y += 2
		case "ne":
			y -= 1
			x += 1
		case "sw":
			y += 1
			x -= 1
		case "nw":
			y -= 1
			x -= 1
		case "se":
			y += 1
			x += 1
		default:
			panic("wtf")
		}
		if d := dist(y, x); d > max {
			max = d
		}
	}
	fmt.Println("part 1:", dist(y, x))
	fmt.Println("part 2:", max)
}
