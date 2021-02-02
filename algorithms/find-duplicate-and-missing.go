// ---------------------------------------
// Find the duplicate and missing numbers
// ---------------------------------------
// Given an array of length n containing numbers from 1,...,n
// find out the number that occurs twice and the missing number
//
// Example:
// Input: [0, 4, 1, 4, 2]
// Output: [4, 3]

package main

import "fmt"

// Time: O(n), Space: O(1)
func findDuplicateAndMissing(arr []int) []int {
	n := len(arr)
	duplicate := -1
	// Duplicated element
	for i := 0; i < n; i++ {
		index := abs(arr[i]) - 1

		if arr[index] < 0 {
			duplicate = abs(arr[i])
		} else {
			arr[index] *= -1
		}
	}

	// Missing element
	missing := -1
	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			missing = i + 1
		}
	}

	return []int{duplicate, missing}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func main() {
	arr := []int{1, 2, 2, 4}
	fmt.Println(findDuplicateAndMissing(arr))
}
