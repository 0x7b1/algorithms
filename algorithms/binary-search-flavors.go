package main

import "fmt"

func basicBinarySearchClosedInterval(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func leftBoundBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) / 2
		if target == nums[mid] {
			right = mid
		} else if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	fmt.Println(basicBinarySearchClosedInterval([]int{1, 2, 3, 4, 5, 6, 7}, 2))
	fmt.Println(leftBoundBinarySearch([]int{1, 2, 2, 2, 3}, 2))
}
