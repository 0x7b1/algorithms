package main

import "fmt"

// O(N*2^N)
func subsets1(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	n := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	res := subsets1(nums)

	size := len(res)
	for i := 0; i < size; i++ {
		res = append(res, res[i])
		res[len(res)-1] = append(res[len(res)-1], n)
	}

	return res
}

func subsets(nums []int) [][]int {
	return nil
}

func backtrack(nums []int, start int, track []int) {
}

func main() {
	fmt.Println(subsets1([]int{1, 2, 3}))
}
