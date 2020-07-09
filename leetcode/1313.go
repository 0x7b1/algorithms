package main

import "fmt"

func decompressRLElist(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			result = append(result, nums[i+1])
		}
	}

	return result
}

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}
