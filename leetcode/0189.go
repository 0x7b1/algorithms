package main

import "fmt"

func rotate(nums []int, k int)  {
	l := len(nums)
	k %= l
	copy(nums, append(nums[l - k:], nums[:l-k]...))
}

func main() {
	arr := []int{-1,-100,3,99}
	rotate(arr, 2)
	fmt.Println(arr)
}
