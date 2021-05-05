package main

import "fmt"

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	avg := 0

	for _, n := range arr {
		sum += n
	}

	avg = sum / 3

	part := 0
	count := 0
	for _, n := range arr {
		part += n
		if part == avg {
			count++
			part = 0
		}
	}

	return count >= 3 && sum % 3 == 0
}

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
}
