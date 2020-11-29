//advent of code 2017, day 23, part 1
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	prog := make([][]string, 0)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		cmd := make([]string, 0)
		for _, token := range strings.Split(strings.TrimSpace(line), " ") {
			cmd = append(cmd, strings.TrimSpace(token))
		}
		prog = append(prog, cmd)
	}
	reg := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "mul": 0}
	isReg := func(s string) bool { return len(s) == 1 && s[0] >= 'a' && s[0] <= 'h' }
	for i := 0; i >= 0 && i < len(prog); {
		var err error
		cmd, x, y := prog[i][0], prog[i][1], prog[i][2]
		xv, yv := 0, 0

		if isReg(x) {
			xv = reg[x]
		} else {
			xv, err = strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
		}

		if isReg(y) {
			yv = reg[y]
		} else {
			yv, err = strconv.Atoi(y)
			if err != nil {
				panic(err)
			}
		}
		switch cmd {
		case "jnz":
			if xv != 0 {
				i += yv
			} else {
				i++
			}
		case "mul":
			reg[x] *= yv
			reg["mul"]++
			i++
		case "set":
			reg[x] = yv
			i++
		case "sub":
			reg[x] -= yv
			i++
		default:
			panic(prog[i])
		}
	}
	fmt.Println("part 1:", reg["mul"])
}
