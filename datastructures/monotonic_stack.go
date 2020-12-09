package main

import "fmt"

// Complexity: O(n)
func nextGreaterElement(nums []int) []int {
	var stack []int
	n := len(nums)
	res := make([]int, n)
	var topIdxStack int

	for i := n - 1; i >= 0; i-- {
		topIdxStack = len(stack) - 1
		for len(stack) > 0 && stack[topIdxStack] <= nums[i] {
			stack = stack[:topIdxStack]
			topIdxStack = len(stack) - 1
		}

		if len(stack) > 0 {
			res[i] = stack[topIdxStack] // If we want to have the next greater one
		} else {
			res[i] = -1
		}

		stack = append(stack, nums[i])
	}

	return res
}

func main() {
	fmt.Println(nextGreaterElement([]int{2, 1, 2, 4, 3}))
	fmt.Println(nextGreaterElement([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
