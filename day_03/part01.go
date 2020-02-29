/* advent of code 2017: day 03, part 01 */

package main

import (
    "fmt"
)

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    var i, y, x int

    input := 325489
    for i=1; i*i < input; i+=2 {}
    calc := i*i
    done := false
    for y=i-1; y>=0 && !done; y-- {
        for x=i-1; x>=0 && !done; x-- {
            if (calc == input) {
                done = true
            }
            calc--;
        }
    }
    cy, cx := (i-1)/2, (i-1)/2
    res := abs((y/i) - (cy/i)) + abs((x%i) - (cx%i))
    fmt.Println(res)
}
