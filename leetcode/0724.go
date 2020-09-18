package main

import "fmt"

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	left := 0
	for i, num := range nums {
		if left * 2 + num == sum {
			return i
		}

		left += num
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
}
