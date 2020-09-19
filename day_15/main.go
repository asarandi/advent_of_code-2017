// advent of code 2017, day 15, part 1 and 2
package main

import (
	"fmt"
)

const valueA = 699 //65 sample
const valueB = 124 //8921 sample

func main() {
	res, a, b := 0, uint(valueA), uint(valueB)
	for i := 0; i < 40000000; i++ {
		a = a * 16807 % 2147483647
		b = b * 48271 % 2147483647
		if a&0xffff == b&0xffff {
			res++
		}
	}
	fmt.Println("part 1:", res)

	res, a, b = 0, uint(valueA), uint(valueB)
	haveA, haveB := false, false
	for i := 0; i < 5000000; {
		if !haveA {
			a = a * 16807 % 2147483647
			haveA = a&3 == 0
		}
		if !haveB {
			b = b * 48271 % 2147483647
			haveB = b&7 == 0
		}
		if haveA && haveB {
			haveA, haveB = false, false
			if a&0xffff == b&0xffff {
				res++
			}
			i++
		}
	}
	fmt.Println("part 2:", res)
}
