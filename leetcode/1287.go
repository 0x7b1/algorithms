package main

import "fmt"

func findSpecialInteger(arr []int) int {
	n := len(arr)
	t := n / 4

	for i := 0; i < n-t; i++ {
		if arr[i] == arr[i + t] {
			return arr[i]
		}
	}

	return -1
}

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}
