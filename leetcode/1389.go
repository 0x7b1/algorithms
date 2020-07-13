package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	size := len(index)
	arr := make([]int, size)
	idx := 0
	for i := 0; i < size; i++ {
		idx = index[i]
		copy(arr[idx:], append([]int{nums[i]}, arr[idx:]...))
	}

	return arr
}

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}
