package main

import "fmt"

func runningSum(nums []int) []int {
    for i := 0; i < len(nums) - 1; i++ {
        nums[i+1] += nums[i]
    }

    return nums
}

func main() {
	a := []int{1,2,3}
	fmt.Println(runningSum(a))
}
