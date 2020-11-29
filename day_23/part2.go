//advent of code 2017, day 23, part 2
package main

import (
	"fmt"
)

const (
	from      = 109900
	to        = 126900
	increment = 17
)

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	res := 0
	for i := from; i <= to; i += increment {
		if !isPrime(i) {
			res++
		}
	}
	fmt.Println("part 2:", res)
}
