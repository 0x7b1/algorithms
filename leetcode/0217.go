package main

import "fmt"

func containsDuplicate(nums []int) bool {
	counters := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := counters[nums[i]]; ok {
			return true
		}

		counters[nums[i]] = true
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
