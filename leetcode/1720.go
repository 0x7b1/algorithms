package main

import "fmt"

func decode(encoded []int, first int) []int {
	arr := []int{first}

	for i, n := range encoded {
		next := n ^ arr[i]
		arr = append(arr, next)
	}

	return arr
}

func main() {
	fmt.Println(decode([]int{1, 2, 3}, 1))
	fmt.Println(decode([]int{6, 2, 7, 3}, 4))
}
