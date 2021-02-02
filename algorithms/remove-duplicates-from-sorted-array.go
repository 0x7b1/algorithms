package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow, fast := 0, 1
	for fast < n {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}

		fast++
	}

	return slow + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
