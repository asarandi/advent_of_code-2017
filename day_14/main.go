//advent of code 2017, day 14, part 1 and 2
package main

import (
	"fmt"
)

const puzzleInput = `oundnydw`

//const puzzleInput = `flqrgnkx` //sample

func countIslands(data [][]int) int {
	f := func(data [][]int, y, x, k int) {}
	f = func(data [][]int, y, x, k int) {
		if y < 0 || y >= len(data) {
			return
		}
		if x < 0 || x >= len(data[y]) {
			return
		}
		if data[y][x] != -1 {
			return
		}
		data[y][x] = k
		f(data, y+1, x, k)
		f(data, y-1, x, k)
		f(data, y, x+1, k)
		f(data, y, x-1, k)
	}
	res := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == -1 {
				res++
				f(data, y, x, res)
			}
		}
	}
	return res
}

func main() {
	setBits := 0
	grid := make([][]int, 128)
	for i := 0; i < 128; i++ {
		grid[i] = make([]int, 128)
		hash := knotHash([]byte(fmt.Sprintf("%s-%d", puzzleInput, i)))
		for j, b := range hash {
			for k := 0; k < 8; k++ {
				if b&(1<<uint(7-k)) != 0 {
					setBits++
					grid[i][j*8+k] = -1
				}
			}
		}
	}
	fmt.Println("part 1:", setBits)
	fmt.Println("part 2:", countIslands(grid))
}
