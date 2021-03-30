package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(num int) []string {
	var res []string

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if bits.OnesCount(uint(h * 64 + m)) == num {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return res
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
