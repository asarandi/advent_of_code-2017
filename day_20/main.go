//advent of code 2017, day 20, part 1 and 2
package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type vec3f struct {
	x, y, z float64
}

type particle struct {
	p, v, a  *vec3f
	dcount   float64
	dsum     float64
	daverage float64
	col      bool
}

func (v *vec3f) add(r *vec3f) {
	v.x += r.x
	v.y += r.y
	v.z += r.z
}

func (p particle) distance() float64 {
	return math.Abs(p.p.x) + math.Abs(p.p.z) + math.Abs(p.p.y)
}

func parse(s string) *vec3f {
	f := make([]float64, 3)
	data := strings.Split(s[3:len(s)-1], ",")
	for i := 0; i < 3; i++ {
		val, err := strconv.ParseFloat(data[i], 64)
		if err != nil {
			panic(err)
		}
		f[i] = val
	}
	return &vec3f{x: f[0], y: f[1], z: f[2]}
}

func average(f []float64) (res float64) {
	for _, v := range f {
		res += v
	}
	return res / float64(len(f))
}

var particles1 []particle
var particles2 []particle

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	particles1 = make([]particle, 0)
	particles2 = make([]particle, 0)
	for _, line := range lines {
		v := strings.Split(line, ", ")
		p1 := particle{p: parse(v[0]), v: parse(v[1]), a: parse(v[2])}
		p2 := particle{p: parse(v[0]), v: parse(v[1]), a: parse(v[2])}
		particles1 = append(particles1, p1)
		particles2 = append(particles2, p2)
	}

	scores := map[int]int{}
	for i := 0; i < 999; i++ {
		ci, cv := -1, math.Inf(1)
		for j, p := range particles1 {
			p.dsum += p.distance()
			p.dcount += 1
			p.daverage = p.dsum / p.dcount
			if p.daverage < cv {
				ci, cv = j, p.daverage
			}
			p.v.add(p.a)
			p.p.add(p.v)
		}
		scores[ci]++
	}
	i, v := -1, -1
	for ci, cv := range scores {
		if cv > v {
			v = cv
			i = ci
		}
	}
	log.Println("part 1:", i)

	for i = 0; i < 999; i++ {
		collisions := make(map[vec3f]int)
		for j, p := range particles2 {
			if p.col {
				continue
			}
			k, ok := collisions[*p.p]
			if ok {
				particles2[k].col, particles2[j].col = true, true
			}
			collisions[*p.p] = j
			p.v.add(p.a)
			p.p.add(p.v)
		}
	}
	v = 0
	for _, p := range particles2 {
		if p.col {
			v++
		}
	}
	log.Println("part 2:", len(particles2)-v)
}
