package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	diff := 0
	sort.Ints(arr)
	diff = arr[1] - arr[0]

	for i := 0; i < len(arr)-1; i++ {
		if ((arr[i+1] - arr[i]) ^ diff) != 0 {
			return false
		}

		diff = arr[i+1] - arr[i]
	}

	return true
}

func main() {
	//fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4}))
}
