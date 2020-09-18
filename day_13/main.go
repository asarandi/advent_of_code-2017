package main

import (
    "io/ioutil"
    "fmt"
    "strconv"
    "strings"
)

type Scanner struct {
    Depth uint
    Range uint
    Position uint
    Direction uint
}

func (s *Scanner) String() string {
    return fmt.Sprintf("Scanner{} - Depth: [%d] Range: [%d] Position: [%d] Direction: [%d]",
        s.Depth, s.Range, s.Position, s.Direction)
}

func (s *Scanner) Reset() {
    s.Position, s.Direction = 0, 0
}

func (s *Scanner) Tick() {
    if s.Direction == 0 {
        if s.Position < s.Range-1 {
            s.Position += 1
        }
    }
    if s.Direction == 1 {
        if s.Position > 0 {
            s.Position -= 1
        }
    }
    if s.Position == 0 {
        s.Direction = 0
    }
    if s.Position == s.Range-1 {
        s.Direction = 1
    }
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    scanners := make([]*Scanner, 0)
    var maxDepth uint
    for _, line := range strings.Split(string(data), "\n") {
        split := strings.Split(strings.TrimSpace(line), ": ")
        if len(split) != 2 {
            continue
        }
        dpt, err := strconv.Atoi(split[0])
        if err != nil {
            panic(err)
        }
        rng, err := strconv.Atoi(split[1])
        if err != nil {
            panic(err)
        }
        scanners = append(scanners, &Scanner{Depth: uint(dpt), Range: uint(rng)})
        if uint(dpt) > maxDepth {
            maxDepth = uint(dpt)
        }
    }
    severity := uint(0)
    for i := uint(0); i<=maxDepth; i++ {
        for _, scanner := range scanners {
            if scanner.Depth == i && scanner.Position == 0 {
                severity += scanner.Depth * scanner.Range
            }
            scanner.Tick()
        }
    }
    fmt.Println("part 1:", severity)
}
