package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		e := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > e {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = e
	}
}

func main() {
	arr := []int{4, 3, 2, 1}
	insertionSort(arr)
	fmt.Println(arr)
}
