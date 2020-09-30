package main

import "fmt"

func sumSubArray(arr []int, n int) int {
	var sum int

	for i := 0; i < len(arr); i++ {
		if i+n > len(arr) {
			continue
		}

		cut := arr[i : i+n]
		for _, e := range cut {
			sum += e
		}

	}

	return sum
}

func sumOddLengthSubarrays(arr []int) int {
	var sum int
	for n := 1; n <= len(arr); n += 2 {
		sum += sumSubArray(arr, n)
	}

	return sum
}

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	fmt.Println(sumOddLengthSubarrays([]int{1, 2}))
	fmt.Println(sumOddLengthSubarrays([]int{10, 11, 12}))
}
