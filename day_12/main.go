// advent of code 2017, day 12, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	Id      int
	Links   map[*Node]bool
	Visited bool
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	allNodes := make(map[int]*Node)
	re := regexp.MustCompile(`(\d+)`)
	for _, line := range strings.Split(string(data), "\n") {
		subs := re.FindAllStringSubmatch(line, -1)
		if len(subs) < 2 {
			continue
		}
		var parent int
		for i, s := range subs {
			id, _ := strconv.Atoi(s[0])
			node := &Node{Id: id, Links: make(map[*Node]bool)}
			if val, ok := allNodes[id]; ok {
				node = val
			} else {
				allNodes[id] = node
			}
			if i == 0 {
				parent = id
			} else {
				allNodes[parent].Links[node] = true
			}
		}
	}

	f := func(group int) (numNodes int) {
		val, ok := allNodes[group]
		if !ok || val.Visited {
			return
		}
		list := []int{group}
		for len(list) > 0 {
			id := list[0]
			list = list[1:]
			if allNodes[id].Visited {
				continue
			}
			allNodes[id].Visited = true
			numNodes++
			for link := range allNodes[id].Links {
				if link.Visited {
					continue
				}
				list = append(list, link.Id)
			}
		}
		return
	}

	fmt.Println("part 1:", f(0))
	groupCount := 1
	for key := range allNodes {
		if f(key) > 0 {
			groupCount++
		}
	}
	fmt.Println("part 2:", groupCount)
}
