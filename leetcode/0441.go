package main

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	left, right := 0, n

	for left <= right {
		k := (left + right) / 2
		curr := k * (k + 1) / 2

		if curr == n {
			return k
		} else if curr > n {
			right = k - 1
		} else {
			left = k + 1
		}
	}

	return right
}

func arrangeCoins1(n int) int {
	return int(math.Sqrt(float64(2*n)+0.25) - 0.5)
}

func main() {
	fmt.Println(arrangeCoins(5))
	fmt.Println(arrangeCoins(8))
}
