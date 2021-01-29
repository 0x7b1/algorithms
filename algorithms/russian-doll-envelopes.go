package main

import (
	"fmt"
	"sort"
)

type Pairs [][]int

func (a Pairs) Len() int {
	return len(a)
}

func (a Pairs) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a Pairs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	sort.Sort(Pairs(envelopes))
	fmt.Println(envelopes)
	height := make([]int, n)

	for i := 0; i < n; i++ {
		height[i] = envelopes[i][1]
	}

	return lengthOfLongestIncreasingSubsequence(height)
}

func lengthOfLongestIncreasingSubsequence(nums []int) (piles int) {
	top := make([]int, len(nums))

	for _, num := range nums {
		left, right := 0, piles // TODO: revise pilar's value
		for left < right {
			mid := (left + right) / 2
			if top[mid] >= num {
				right = mid
			} else {
				left = mid + 1
			}

			if left == num {
				piles++
			}

			top[left] = num
		}
	}

	return
}

func main() {
	arr := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	fmt.Println(maxEnvelopes(arr))
}
