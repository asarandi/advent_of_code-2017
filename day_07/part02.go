/* advent of code 2017: day 07, part 02 */
package main

import (
	"fmt"
	_"log"
	"strings"
	"strconv"
	"io/ioutil"
	"regexp"
)

var childParent map[string]string
var parentChildren map[string][]string
var nodeWeight map[string]int
var root string

func sum(p string) int {
	res := nodeWeight[p]
	k := len(parentChildren[p])
	for i:=0; i<k; i++ {
		c := parentChildren[p][i]
		res += sum(c)
	}
	return res
}

func prnt(p string) {
	ps := sum(p)
	fmt.Printf("p[%8s] w[%5d] ", p, ps)
	k := len(parentChildren[p])
	for i:=0; i<k; i++ {
		c := parentChildren[p][i]
		cs := sum(c)
		fmt.Printf("c[%8s] w[%5d] ", c, cs)
	}
	fmt.Printf("\n")
}

func balanced(p string) bool {
	k := len(parentChildren[p])
	res := true
	for i:=0; i+1<k; i++ {
		a := parentChildren[p][i]
		b := parentChildren[p][i+1]
		if sum(a) != sum(b) {
			prnt(p)
			if sum(a) > sum(b) {
				balanced(a)
			} else {
				balanced(b)
			}
			res = false
			break
		}
	}
	return res
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil { panic(err) }
	split1 := strings.Split(strings.Trim(string(content), " \t\r\n\v\f"), "\n")
	re := regexp.MustCompile("(\\S+)\\s+\\((\\d+)\\)\\s?\\-?\\>?\\s?(.*)?")
	if err != nil { panic(err) }
	childParent = make(map[string]string)
	parentChildren = make(map[string][]string)
	nodeWeight = make(map[string]int)
	for _, line := range split1 {
		rss := re.FindStringSubmatch(line)
		p := rss[1]
		weight, _ := strconv.Atoi(rss[2])
		nodeWeight[p] = weight
		if len(rss[3]) > 0 {
			_, ok := parentChildren[p]
			if !ok {
				parentChildren[p] = make([]string,0)
			}
			split2 := strings.Split(rss[3], ", ")
			for _, child := range split2 {
				childParent[child] = p
				parentChildren[p] = append(parentChildren[p], child)
			}
		}
	}
	for _,v := range childParent {
		_, ok := childParent[v]
		if !ok {
			root = v
			fmt.Println("part 1:", root)
			break
		}
	}
	prnt("tulwp")
	fmt.Println(nodeWeight["tulwp"])
//	balanced(root)
}
