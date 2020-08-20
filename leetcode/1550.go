package main

import (
	"fmt"
)

func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for _, n := range arr {
		if n%2 != 0 {
			count++
		} else {
			count = 0
		}

		if count == 3 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(threeConsecutiveOdds([]int{424, 915, 193, 591, 923}))
	// fmt.Println(threeConsecutiveOdds([]int{1, 2, 1, 1}))
	// fmt.Println(threeConsecutiveOdds([]int{1, 1, 1}))
	// fmt.Println(threeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12}))
	// fmt.Println(threeConsecutiveOdds([]int{2, 6, 4, 1}))
}
