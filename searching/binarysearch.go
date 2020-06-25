package main

import "fmt"

func search(arr []int, n int) bool {
	middle := len(arr) / 2
	if arr[middle] == n {
		return true
	} else if arr[middle] < 2 {
		return search(arr[:middle], n)
	} else {
		return search(arr[middle + 1:], n)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	found := search(arr, 22)
	fmt.Println(found)
}
