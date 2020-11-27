/* advent of code 2017: day 01, part 01 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Trim(string(content), "\n")
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == s[(i+1)%len(s)] {
			res += int(s[i] - '0')
		}
	}
	fmt.Println(res)
}
