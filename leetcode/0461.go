package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance2(x int, y int) int {
	diff := x ^ y
	count := 0

	for diff > 0 {
		count += diff & 1
		diff >>= 1
	}

	return count
}

func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
