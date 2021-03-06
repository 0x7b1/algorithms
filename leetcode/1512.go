package main

import "fmt"

// num[i] == num[j] && i < j
func numIdenticalPairs(nums []int) int {
	pairs := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				pairs++
			}
		}
	}

	return pairs
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))
}
