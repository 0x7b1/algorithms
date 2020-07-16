package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	sortedHeights := make([]int, len(heights))
	copy(sortedHeights, heights)

	sort.Ints(sortedHeights)

	swaps := 0
	for i, n := range heights {
		if sortedHeights[i] != n {
			swaps++
		}
	}

	return swaps
}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}
