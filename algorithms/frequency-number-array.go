package main

import "fmt"

func leftBound(arr []int, n int) int {
	l, r := 0, len(arr)

	for l < r {
		m := l + (r-l)/2
		if arr[m] == n {
			r = m
		} else if arr[m] < n {
			l = m + 1
		} else {
			r = m
		}
	}

	return l
}

func rightBound(arr []int, n int) int {
	l, r := 0, len(arr)

	for l < r {
		m := l + (r-l)/2
		if arr[m] == n {
			l = m + 1
		} else if arr[m] < n {
			l = m + 1
		} else {
			r = m
		}
	}

	return l - 1
}

func getFrequency(arr []int, n int) int {
	return rightBound(arr, n) - leftBound(arr, n) + 1
}

func main() {
	fmt.Println(getFrequency([]int{1, 1, 2, 2, 2, 2, 3}, 2))
	fmt.Println(getFrequency([]int{1, 1, 2, 2, 2, 2, 3}, 1))
	fmt.Println(getFrequency([]int{1, 1, 2, 2, 2, 2, 3}, 3))
}
