package main

import "fmt"

func rotate2(nums []int, k int)  {
	l := len(nums)
	k %= l
	copy(nums, append(nums[l - k:], nums[:l-k]...))
}

func rotate(nums []int, k int)  {
	newnums := make([]int, len(nums))
	for i, v := range nums {
		newnums[(i+k)%len(nums)] = v
	}

	copy(nums, newnums)
}

func main() {
	arr := []int{-1,-100,3,99}
	rotate(arr, 2)
	fmt.Println(arr)
}
