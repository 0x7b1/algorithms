package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const maxbit = (1 << 7)
const digit = 5

func radixSort(data []uint8) {
	buff := bytes.NewBuffer(nil)
	ds := make([][]byte, len(data))

	for i, e := range data {
		binary.Write(buff, binary.LittleEndian, e^maxbit)
		b := make([]byte, digit)
		buff.Read(b)
		ds[i] = b
	}

	fmt.Println("->", ds)

	countingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			countingSort[b[i]] = append(countingSort[b[i]], b)
		}
		j := 0
		for k, bs := range countingSort {
			copy(ds[j:], bs)
			j += len(bs)
			countingSort[k] = bs[:0]
		}
	}

	var w uint8
	for i, b := range ds {
		buff.Write(b)
		binary.Read(buff, binary.LittleEndian, &w)
		data[i] = w ^ maxbit
	}
}

func main() {
	arr := []uint8{60, 10, 4, 62, 121}
	radixSort(arr)
	fmt.Println(arr, maxbit)
}
