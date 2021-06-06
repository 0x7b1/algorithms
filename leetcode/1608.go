package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		n := len(nums) - i
		if n <= num && (i - 1 < 0 || n > nums[i-1]) {
			return n
		}
	}

 	return -1
}

func main() {
	fmt.Println(specialArray([]int{3, 5}))
}
