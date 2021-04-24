package main

import (
	"fmt"
)

func findErrorNums(nums []int) []int {
	diff := 0
	squareDiff := 0

	for i := 0; i < len(nums); i++ {
		diff += (i + 1) - nums[i]
		squareDiff += (i+1)*(i+1) - nums[i]*nums[i]
	}

	sum := squareDiff / diff
	return []int{(sum - diff) / 2, (sum + diff) / 2}
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
	fmt.Println(findErrorNums([]int{3, 2, 2}))
	fmt.Println(findErrorNums([]int{3, 2, 3, 4, 6, 5}))
}
