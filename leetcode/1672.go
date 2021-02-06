package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	max := 0

	for _, nrs := range accounts {
		sum := 0
		for _, n := range nrs {
			sum += n
		}

		if sum > max {
			max = sum
		}
	}

	return max
}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}}))
}
