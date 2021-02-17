package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	d := make(map[int]int)

	for i, num := range nums {
		if idx, ok := d[target-num]; ok {
			return []int{idx, i}
		}

		d[num] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
