//adevent of code 2017, day 21, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	_ "math"
	"strings"
)

var (
	square       = [][]uint{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}}
	two2three    map[uint]uint
	three2four   map[uint]uint
	rotateTwo    = map[uint]uint{0: 1, 1: 3, 2: 0, 3: 2}
	rowFlipTwo   = map[uint]uint{0: 2, 1: 3, 2: 0, 3: 1}
	colFlipTwo   = map[uint]uint{0: 1, 1: 0, 2: 3, 3: 2}
	rotateThree  = map[uint]uint{0: 2, 1: 5, 2: 8, 3: 1, 4: 4, 5: 7, 6: 0, 7: 3, 8: 6}
	rowFlipThree = map[uint]uint{0: 6, 1: 7, 2: 8, 3: 3, 4: 4, 5: 5, 6: 0, 7: 1, 8: 2}
	colFlipThree = map[uint]uint{0: 2, 1: 1, 2: 0, 3: 5, 4: 4, 5: 3, 6: 8, 7: 7, 8: 6}
)

func toUint(in string) (out uint) {
	for i := 0; i < len(in); i++ {
		if in[i] == '.' || in[i] == '#' {
			out <<= 1
			if in[i] == '#' {
				out |= 1
			}
		}
	}
	return
}

func transform(in uint, moves map[uint]uint) (out uint) {
	for from, to := range moves {
		out |= ((in >> from) & 1) << to
	}
	return
}

func permuteTwos() {
	added := make(map[uint]uint)
	for from, to := range two2three {
		for _, input := range []uint{
			from,
			transform(from, rowFlipTwo),
			transform(from, colFlipTwo),
		} {
			for i := 0; i < 4; i++ {
				rotated := transform(input, rotateTwo)
				added[rotated] = to
				input = rotated
			}
		}
	}
	for from, to := range added {
		two2three[from] = to
	}
}

func permuteThrees() {
	added := make(map[uint]uint)
	for from, to := range three2four {
		for _, input := range []uint{
			from,
			transform(from, rowFlipThree),
			transform(from, colFlipThree),
		} {
			for i := 0; i < 4; i++ {
				rotated := transform(input, rotateThree)
				added[rotated] = to
				input = rotated
			}
		}
	}
	for from, to := range added {
		three2four[from] = to
	}
}

func transfer(b, a int, pattern map[uint]uint) [][]uint {
	n := len(square) / b
	res := make([][]uint, n*a)
	for y := 0; y < n*a; y++ {
		res[y] = make([]uint, n*a)
	}
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			from := uint(0)
			for i := 0; i < b; i++ {
				for j := 0; j < b; j++ {
					from <<= 1
					from |= square[y*b+i][x*b+j]
				}
			}
			to, ok := pattern[from]
			if !ok {
				panic("not found")
			}
			for i := 0; i < a; i++ {
				for j := 0; j < a; j++ {
					bit := (to >> ((a*a - 1) - (i*a + j))) & 1
					res[y*a+i][x*a+j] = bit
				}
			}
		}
	}
	return res
}

func tick() {
	var newSquare [][]uint
	if len(square)%2 == 0 {
		newSquare = transfer(2, 3, two2three)
	} else {
		newSquare = transfer(3, 4, three2four)
	}
	square = newSquare
}

func count() (res uint) {
	for i := 0; i < len(square); i++ {
		for j := 0; j < len(square); j++ {
			res += square[i][j]
		}
	}
	return
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	two2three = make(map[uint]uint)
	three2four = make(map[uint]uint)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		split := strings.Split(line, " => ")
		from, to := toUint(split[0]), toUint(split[1])
		if len(split[0]) == 5 {
			two2three[from] = to
		} else {
			three2four[from] = to
		}
	}
	permuteTwos()
	permuteThrees()

	for i := 0; i < 5; i++ {
		tick()
	}
	fmt.Println("part 1:", count())
	for i := 0; i < 13; i++ {
		tick()
	}
	fmt.Println("part 2:", count())
}
