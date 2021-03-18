package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i > 1; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))
	fmt.Println(largestPerimeter([]int{1, 2, 1}))
	fmt.Println(largestPerimeter([]int{3, 2, 3, 4}))
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))

	fmt.Println("HELLO WORLD")
}
