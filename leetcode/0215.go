package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)

	return nums[len(nums) - k]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
