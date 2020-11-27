/* advent of code 2017: day 01, part 02 */
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
	j := len(s)
	for i := 0; i < j; i++ {
		if s[i] == s[(i+(j>>1))%j] {
			res += int(s[i] - '0')
		}
	}
	fmt.Println(res)
}
