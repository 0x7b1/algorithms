package main

import "fmt"

func peakIndexInMountainArray2(A []int) int {
	max, idx := A[0], 0

	for i := 1; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
			idx = i
		}
	}

	return idx
}

func peakIndexInMountainArray(A []int) int {
	lo, mid, hi := 1, 0, len(A)-1

	for lo <= hi {
		mid = (lo + hi) / 2
		if A[mid] > A[mid-1] && A[mid] > A[mid+1] {
			break
		} else if A[mid] < A[mid+1] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return mid
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
}
