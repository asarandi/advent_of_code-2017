package main

import "bytes"

func knotHash(input []byte) []byte {
	const listSize = 256
	var i, j, k, pos, skip uint

	lengths := bytes.Join([][]byte{input, {17, 31, 73, 47, 23}}, nil)
	data := make([]byte, listSize)
	for i = 0; i < listSize; i++ {
		data[i] = byte(i)
	}
	for i = 0; i < 64; i++ {
		for _, b := range lengths {
			k := uint(b)
			func(data []byte, i, j uint) {
				for i < j {
					temp := data[i%listSize]
					data[i%listSize] = data[j%listSize]
					data[j%listSize] = temp
					i += 1
					j -= 1
				}
			}(data, pos, pos+k-1)
			pos += k + skip
			skip += 1
		}
	}
	res := make([]byte, 16)
	for i = 0; i < 16; i++ {
		for j, k = 0, 0; j < 16; j++ {
			k ^= uint(data[i*16+j])
		}
		res[i] = byte(k)
	}
	return res
}
