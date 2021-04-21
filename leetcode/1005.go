package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	minIdx := 0
	for i := 0; i < K; i++ {
		A[minIdx] = -A[minIdx]
		if A[minIdx+1] < A[minIdx] {
			minIdx++
		}
	}
	
	sum := 0
	for _, n := range A {
		sum += n
	}

	return sum
}

func main() {
	fmt.Println(largestSumAfterKNegations([]int{4, 2, 3}, 1))
	fmt.Println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	fmt.Println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}
