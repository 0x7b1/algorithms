package main

import (
	"fmt"
)

func replaceElements(arr []int) []int {
	j, i, l := -1, -2, len(arr)
	for l--; l >= 0; l-- {
		if j > i {
			i = arr[l]
			arr[l] = j
		} else {
			j = arr[l]
			arr[l] = i
		}
	}

	return arr
}

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}
