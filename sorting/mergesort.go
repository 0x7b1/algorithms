package main

import "fmt"

func merge(a, b []int) []int {
	i, j := 0, 0
	var c []int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	for i < len(a) {
		c = append(c, a[i])
		i++
	}

	for j < len(b) {
		c = append(c, b[j])
		j++
	}

	return c
}

func mergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	middle := len(a) / 2
	left := mergeSort(a[:middle])
	right := mergeSort(a[middle:])

	return merge(left, right)
}

func main() {
	arr := []int{5, 4, 33, 2, 1}
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)
}
