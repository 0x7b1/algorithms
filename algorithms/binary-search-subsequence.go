package main

import "fmt"

func leftBound(arr []int, tar int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if tar > arr[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// Time: O(MlogN)
func isSubsequence(s, t string) bool {
	// Store indices of each character in a dictionary
	index := make([][]int, 256)

	for i, c := range t {
		if index[c] == nil {
			index[c] = make([]int, 0)
		}

		index[c] = append(index[c], i)
	}

	// look for subsequence
	j := 0
	for _, c := range s {
		if index[c] == nil {
			return false
		}

		pos := leftBound(index[c], j)
		if pos == len(index[c]) {
			return false
		}

		j = index[c][pos] + 1
	}

	return true
}

func main() {
	fmt.Println(isSubsequence("abc", "cacbhbc"))
	fmt.Println(isSubsequence("axc", "cacbhbc"))
}
