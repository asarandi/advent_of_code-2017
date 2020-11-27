/* advent of code 2017: day 02, part 02 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	rd := bufio.NewReader(fp)
	res := 0
	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		str = strings.Trim(str, " \t\n\r\v\f")
		split := strings.Split(str, "\t")
		array := make([]int, len(split))
		for i, e := range split {
			val, _ := strconv.Atoi(e)
			array[i] = val
		}
		done := false
		for i := 0; i+1 < len(array) && !done; i++ {
			for j := i + 1; j < len(array) && !done; j++ {
				var big, small int
				big, small = array[i], array[j]
				if small > big {
					big, small = array[j], array[i]
				}
				if big%small == 0 {
					done = true
					res += big / small
				}
			}
		}
	}
	fmt.Println(res)
}
