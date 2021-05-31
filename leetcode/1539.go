package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	nrs := make(map[int]bool)

	for _, n := range arr {
		nrs[n] = true
	}

	counter := 1
	for {
		if !nrs[counter] {
			k--
		}

		if k == 0 {
			return counter
		}

		counter++
	}
}

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
}
