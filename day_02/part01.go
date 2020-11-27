/* advent of code 2017: day 02, part 01 */
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
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[0])
		for _, e := range split {
			val, _ := strconv.Atoi(e)
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}
		}
		res += max - min
	}
	fmt.Println(res)
}
