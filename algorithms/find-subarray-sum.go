package main

import "fmt"

// Time O(n^2)
func subarraySumBruteforce(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := 0
			for _, n := range nums[i : j+1] {
				sum += n
			}

			if sum == k {
				res++
			}
		}
	}

	return res
}

// Time O(n^2)
func subarraySumWithPrefixSum(nums []int, k int) int {
	res := 0
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if preSum[j+1]-preSum[i] == k {
				res++
			}
		}
	}

	return res
}

// Time O(n)
func subarraySumMemoizedWithPrefix(nums []int, k int) int {
	n := len(nums)
	preSum := make(map[int]int) // presum: frequency
	preSum[0] = 1

	res := 0
	isum := 0

	for i := 0; i < n; i++ {
		isum += nums[i]
		jsum := isum - k

		if _, ok := preSum[jsum]; ok {
			res++
			//res += preSum[jsum]
		}

		preSum[isum]++
	}

	return res
}

func main() {
	arr := []int{3, 5, 2, -2, 4, 1}
	k := 5
	fmt.Println(subarraySumBruteforce(arr, k))
	fmt.Println(subarraySumWithPrefixSum(arr, k))
	fmt.Println(subarraySumMemoizedWithPrefix(arr, k))
}
