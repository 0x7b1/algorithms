package main

import (
	"fmt"
	"sort"
)

func sortedSquares(A []int) []int {
	for i, n := range A {
		A[i] *= n
	}

	sort.Ints(A)

	return A
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
