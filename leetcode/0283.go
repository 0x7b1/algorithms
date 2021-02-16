package main

import "fmt"

func moveZeroes(nums []int) {
	// 0 1 0 3 12
	// 1 0 0 3 12
	// 1 3 0 0 12
	// 1 3 12 0 0

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}

		fast++
	}
}

func main() {
	//arr := []int{2, 1}
	arr := []int{0, 1, 0, 3, 12}
	//arr := []int{1, 0, 1}
	//arr := []int{1, 2, 3, 1}
	//arr := []int{1, 0}
	moveZeroes(arr)
	fmt.Println(arr)
}
