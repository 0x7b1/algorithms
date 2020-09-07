package main

import "fmt"

func minCostToMoveChips(position []int) int {
	ones, zeros := 0, 0
	for i, _ := range position {
		if i % 2 == 0 {
			zeros++
		} else {
			ones++
		}
	}


	if zeros < ones {
		return zeros
	}

	return ones
}

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
	//fmt.Println(minCostToMoveChips([]int{1,1000000000}))
}
