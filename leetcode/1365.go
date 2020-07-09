package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	smaller := make([]int, len(nums))
	var count int

	for i, a := range nums {
		for _, b := range nums {
			if b < a {
				count++
			}
		}

		smaller[i] = count
		count = 0
	}

	return smaller
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent([]int{6,5,4,8}))
}
