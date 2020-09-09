package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	res := [][]int{}
	min := int(^uint(0) >> 1)

	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < min {
			min = diff
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff == min {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}

	return res
}

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
	fmt.Println(minimumAbsDifference([]int{40, 11, 26, 27, -20}))
}
