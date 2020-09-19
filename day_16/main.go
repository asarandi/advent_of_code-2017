//advent of code 2017, day 16, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const numPrograms = 16

var (
	atoi = make(map[byte]int)
	itoa = make(map[int]byte)
)

func sync(f bool) {
	if f {
		for k, v := range itoa {
			atoi[v] = k
		}
	} else {
		for k, v := range atoi {
			itoa[v] = k
		}
	}
}

func spin(n int) {
	for k, v := range atoi {
		if v < len(atoi)-n {
			atoi[k] += n
		} else {
			atoi[k] -= len(atoi) - n
		}
	}
	sync(false)
}

func exchange(i, j int) {
	itoa[i], itoa[j] = itoa[j], itoa[i]
	sync(true)
}

func partner(a, b byte) {
	atoi[a], atoi[b] = atoi[b], atoi[a]
	sync(false)
}

func dance(moves []string) {
	for _, move := range moves {
		switch move[0] {
		case 's':
			n, err := strconv.Atoi(move[1:])
			if err != nil {
				panic(err)
			}
			spin(n)
		case 'x':
			xchg := strings.Split(move[1:], "/")
			a, err := strconv.Atoi(xchg[0])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(xchg[1])
			if err != nil {
				panic(err)
			}
			exchange(a, b)
		case 'p':
			partner(move[1], move[3])
		default:
			panic("wtf")
		}
	}
}

func str() string {
	s := make([]byte, numPrograms)
	for i := 0; i < numPrograms; i++ {
		s[i] = itoa[i]
	}
	return string(s)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < numPrograms; i++ {
		atoi[byte('a'+i)] = i
		itoa[i] = byte('a' + i)
	}
	moves := strings.Split(strings.TrimSpace(string(data)), ",")
	seenAtoi := make(map[string]int)
	seenItoa := make(map[int]string)
	for i := 0; ; i++ {
		seenAtoi[str()], seenItoa[i] = i, str()
		dance(moves)
		if _, ok := seenAtoi[str()]; ok {
			break
		}
	}
	fmt.Println("part 1:", seenItoa[1])
	fmt.Println("part 2:", seenItoa[1000000000%len(seenItoa)])
}
