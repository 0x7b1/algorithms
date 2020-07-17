package main

import "fmt"

func search(arr []int, n int) bool {
	middle := len(arr) / 2
	if arr[middle] == n {
		return true
	} else if arr[middle] < 2 {
		return search(arr[:middle], n)
	} else {
		return search(arr[middle+1:], n)
	}
}

// countOnes takes as an input an array of 1s and 0s
// where 1s are always in front of the array
// and the task is to count the number of 1s
func countOnes(arr []int) int {
	low, hi, mid := 0, len(arr) - 1, 0

	for low < hi {
		mid = low + (hi - low) / 2
		if arr[mid] == 1 {
			low = mid + 1
		} else {
			hi = mid
		}
	}

	return low
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	found := search(arr, 4)
	fmt.Println(found)

	fmt.Println(countOnes([]int{1, 1, 1, 0, 0, 0}))
	fmt.Println(countOnes([]int{1, 0, 0, 0, 0, 0}))
	fmt.Println(countOnes([]int{1, 1, 1, 1, 1, 1, 1, 0}))
}
