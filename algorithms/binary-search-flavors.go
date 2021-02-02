package main

import "fmt"

func basicClosedInterval(nums []int, target int) int {
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

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
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

func leftBoyndSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid - 1
		}
	}

	if left >= len(nums) || nums[left] != target {
		return -1
	}

	return left
}

func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left - 1
}

func rightBoundSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if right < 0 || nums[right] != target {
		return - 1
	}

	return right
}

func main() {
	fmt.Println(basicClosedInterval([]int{1, 2, 2, 2, 3}, 4))
	fmt.Println(leftBound([]int{1, 2, 2, 2, 3}, 4))
	fmt.Println(leftBoyndSearch([]int{1, 2, 2, 2, 3}, 4))
	fmt.Println(rightBound([]int{1, 2, 2, 2, 3}, 4))
	fmt.Println(rightBoundSearch([]int{1, 2, 2, 2, 3}, 4))
}
