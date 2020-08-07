package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := countBits(arr[i]), countBits(arr[j])
		if x == y {
			return arr[i] < arr[j]
		} else {
			return x < y
		}
	})
	return arr
}

func countBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}

	return count
}

func main() {
	fmt.Println(countBits(8742))
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}
