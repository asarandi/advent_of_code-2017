//advent of code 2017, day 18, part 1
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
	program := make([][]string, 0)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		cmd := make([]string, 0)
		for _, token := range strings.Split(strings.TrimSpace(line), " ") {
			cmd = append(cmd, strings.TrimSpace(token))
		}
		program = append(program, cmd)
	}
	registers := make(map[string]int)
	for i := 0; i >= 0 && i < len(program); {
		cmd := program[i]
		key, value := cmd[1], 0
		if len(cmd) == 3 {
			value, err = strconv.Atoi(cmd[2])
			if err != nil {
				value = registers[cmd[2]]
			}
		}
		switch program[i][0] {
		case "add":
			registers[key] += value
			i++
		case "jgz":
			if registers[key] > 0 {
				i += value
			} else {
				i++
			}
		case "mod":
			registers[key] %= value
			i++
		case "mul":
			registers[key] *= value
			i++
		case "rcv":
			if registers[key] != 0 {
				i = -1
			} else {
				i++
			}
		case "set":
			registers[key] = value
			i++
		case "snd":
			registers["snd"] = registers[key]
			i++
		default:
			panic(program[i])
		}
	}
	fmt.Println("part 1:", registers["snd"])
}
