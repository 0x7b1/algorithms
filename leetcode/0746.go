package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	f1, f2 := 0, 0

	for i := len(cost) - 1; i >= 0; i-- {
		f0 := cost[i]

		if f1 < f2 {
			f0 += f1
		} else {
			f0 += f2
		}

		f1, f2 = f0, f1
	}

	if f1 < f2 {
		return f1
	}

	return f2
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
