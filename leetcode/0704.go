package main

import "fmt"

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)

	for lo < hi {
		mid := int(uint(lo + hi) >> 1)
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if hi < len(nums)&& nums[lo] == target {
		return lo
	} else {
		return -1
	}
}

func main() {
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6, 7}, 5))
	fmt.Println(search([]int{-1, 2, 3, 4, 5, 6, 7}, -1))
	fmt.Println(search([]int{2, 5}, 5))
	fmt.Println(search([]int{2, 5}, 22))
}
