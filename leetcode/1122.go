package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Set := make(map[int]bool)
	for _, n := range arr2 {
		arr2Set[n] = true
	}

	arr1Count := make(map[int]int)
	leftOvers := make([]int, 0)
	for _, n := range arr1 {
		if !arr2Set[n] {
			leftOvers = append(leftOvers, n)
		} else {
			arr1Count[n]++
		}
	}

	sort.Ints(leftOvers)

	res := make([]int, 0)
	for _, n := range arr2 {
		for i := 0; i < arr1Count[n]; i++ {
			res = append(res, n)
		}
	}

	res = append(res, leftOvers...)

	return res
}

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}
