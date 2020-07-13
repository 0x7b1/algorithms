package main

import (
	"fmt"
	"math"
)

func findNumbers(nums []int) int {
	evens := 0
	for _, num := range nums {
		if (int(math.Log10(float64(num))) + 1) & 1 == 0 {
			evens++
		}
	}

	return evens
}

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
}
