/* advent of code 2017: day 07, part 01 */
package main

import (
	"fmt"
	"io/ioutil"
	_ "log"
	"regexp"
	_ "strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	split1 := strings.Split(strings.Trim(string(content), " \t\r\n\v\f"), "\n")
	re := regexp.MustCompile("(\\S+)\\s+\\((\\d+)\\)\\s?\\-?\\>?\\s?(.*)?")
	if err != nil {
		panic(err)
	}
	childParent := make(map[string]string)
	for _, line := range split1 {
		rss := re.FindStringSubmatch(line)
		if len(rss[3]) > 0 {
			split2 := strings.Split(rss[3], ", ")
			for _, child := range split2 {
				childParent[child] = rss[1]
			}
		}
	}
	for _, v := range childParent {
		_, ok := childParent[v]
		if !ok {
			fmt.Println("part 1:", v)
			break
		}
	}
}
