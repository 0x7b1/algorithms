package main

import (
	"fmt"
	"github.com/pkg/math"
)

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = math.Max(dp[i], dp[j] + 1)
			}
		}
	}

	return math.MaxIntN(dp...)
}

func lengthOfLISBST(nums []int) int {
	top := make([]int, len(nums))
	piles := 0
	for i := 0; i < len(nums); i++ {
		poker := nums[i]
		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == piles {
			piles++
		}

		top[left] = poker
	}

	// The number of piles represent the length of LIS
	return piles
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLISBST([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
