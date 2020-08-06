package main

import "fmt"

func findDuplicates(nums []int) []int {
	numsCount := make(map[int]int)
	for _, n := range nums {
		numsCount[n]++
	}

	twice := []int{}
	for k, v := range numsCount {
		if v == 2 {
			twice = append(twice, k)
		}
	}

	return twice
}

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
