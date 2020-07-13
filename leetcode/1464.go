package main

import "fmt"

func maxProduct(nums []int) int {
	var max1, max2 int

	for _, num := range nums {
		if num > max1 {
			max1, max2 = num, max1
		} else if num > max2 {
			max2 = num
		}
	}

	return (max1 - 1) * (max2 - 1)
}

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
	fmt.Println(maxProduct([]int{1, 5, 4, 5}))
	fmt.Println(maxProduct([]int{3, 7}))
}
