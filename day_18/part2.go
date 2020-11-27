//advent of code 2017, day 18, part 2
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

	registers := make([]map[string]int, 2)
	queues := make([][]int, 2)
	for i := 0; i < 2; i++ {
		registers[i] = make(map[string]int)
		queues[i] = make([]int, 0)
	}
	registers[0]["p"], registers[1]["p"] = 0, 1

	step := func(p int) {
		i := "index"
		if registers[p][i] < 0 || registers[p][i] >= len(program) {
			registers[p]["done"] = 1
			return
		}
		registers[p]["done"] = 0
		cmd := program[registers[p][i]]
		op, key, value := cmd[0], cmd[1], 0
		if len(cmd) == 3 {
			value, err = strconv.Atoi(cmd[2])
			if err != nil {
				value = registers[p][cmd[2]]
			}
		}
		switch op {
		case "add":
			registers[p][key] += value
			registers[p][i]++
		case "jgz":
			if registers[p][key] > 0 {
				registers[p][i] += value
			} else if key == "1" { // XXX input line 34 .. i assumed all first operands were registers - wrong, took a while to notice
				registers[p][i] += value
			} else {
				registers[p][i]++
			}
		case "mod":
			registers[p][key] %= value
			registers[p][i]++
		case "mul":
			registers[p][key] *= value
			registers[p][i]++
		case "rcv":
			q := (p + 1) % 2 //the other program
			if len(queues[q]) == 0 {
				registers[p]["done"] = 1
				return
			}
			value = queues[q][0]
			queues[q] = queues[q][1:]
			registers[p]["num_recv"]++
			registers[p][key] = value
			registers[p][i]++
		case "set":
			registers[p][key] = value
			registers[p][i]++
		case "snd":
			queues[p] = append(queues[p], registers[p][key])
			registers[p][i]++
			registers[p]["num_sent"]++
		default:
			panic("wtf")
		}
		return
	}

	isDone := func() bool {
		return registers[0]["done"] == 1 && registers[1]["done"] == 1
	}

	for !isDone() {
		for j := 0; j < 2; j++ {
			//            cmd := program[registers[j]["index"]]
			//            fmt.Printf("%2d %s \t%#v\n", registers[j]["index"], cmd, registers[j])
			step(j)
		}
		fmt.Println()
	}

	fmt.Println("part 2:", registers[1]["num_sent"])
}
