//advent of code 2017, day 19, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	n, m := 200, 200 //input size
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	split := strings.Split(string(data), "\n")
	grid := make([][]byte, 0)
	for i := 0; i < n; i++ {
		grid = append(grid, []byte(split[i]))
	}
	d, y, x := 0, 0, 0
	res, seen, steps := "", map[int]bool{}, 0
	for x = 0; grid[0][x] != '|'; x++ {
	}
loop:
	for {
		seen[y*m+x], steps = true, steps+1
		durl := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		c := grid[y][x]
		switch c {
		case '|', '-':
			y, x = y+durl[d][0], x+durl[d][1]
		case '+':
			for i := 0; i < 4; i++ {
				yi, xi := y+durl[i][0], x+durl[i][1]
				if seen[yi*m+xi] {
					continue
				}
				if yi < 0 || yi >= n {
					continue
				}
				if xi < 0 || xi >= m {
					continue
				}
				ci := grid[yi][xi]
				if ci != ' ' {
					d, y, x = i, yi, xi
					break
				}
			}
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M':
			res += string(c)
			y, x = y+durl[d][0], x+durl[d][1]
		case 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			res += string(c)
			y, x = y+durl[d][0], x+durl[d][1]
		default:
			break loop
		}
	}
	fmt.Println("part 1:", res)
	fmt.Println("part 2:", steps-1)
}
