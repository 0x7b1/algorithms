package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	numbers := make(map[int]bool, len(nums))

	for _, num := range nums {
		numbers[num] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !numbers[i] {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
