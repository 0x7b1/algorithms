package main

import "fmt"

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums) + 1)

	for i := 0; i < len(nums); i++ {
		preSum[i + 1] = preSum[i] + nums[i]
	}

	return NumArray{preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right + 1] - this.preSum[left]
}

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
