// advent of code 2017, day 10, part 1 and 2

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const listSize = 256

func reverse(data []byte, i, j uint64) {
	for i < j {
		temp := data[i%listSize]
		data[i%listSize] = data[j%listSize]
		data[j%listSize] = temp
		i += 1
		j -= 1
	}
}

func main() {
	var i, j, k, pos, skip uint64
	var err error

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lengths1 := strings.Split(strings.TrimSpace(string(input)), ",")
	data1, data2 := make([]byte, listSize), make([]byte, listSize)
	for i = 0; i < listSize; i++ {
		data1[i], data2[i] = byte(i), byte(i)
	}
	pos, skip = 0, 0
	for _, s := range lengths1 {
		k, err = strconv.ParseUint(s, 10, 64)
		if err != nil {
			panic(err)
		}
		if k > listSize {
			continue
		}
		reverse(data1, pos, pos+k-1)
		pos += k + skip
		skip += 1
	}
	fmt.Println("part 1:", int(data1[0])*int(data1[1]))

	lengths2 := bytes.Join([][]byte{
		[]byte(strings.TrimSpace(string(input))),
		{17, 31, 73, 47, 23},
	}, nil)
	pos, skip = 0, 0
	for i = 0; i < 64; i++ {
		for _, b := range lengths2 {
			k = uint64(b)
			reverse(data2, pos, pos+k-1)
			pos += k + skip
			skip += 1
		}
	}
	dense := make([]byte, 16)
	for i = 0; i < 16; i++ {
		for j, k = 0, 0; j < 16; j++ {
			k ^= uint64(data2[i*16+j])
		}
		dense[i] = byte(k)
	}
	fmt.Println("part 2:", hex.EncodeToString(dense))
}
