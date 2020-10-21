package main

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	var groups [][]int
	tmp := make(map[int][]int)

	for i, size := range groupSizes {
		if _, ok := tmp[size]; !ok {
			tmp[size] = []int{}
		}

		tmp[size] = append(tmp[size], i)

		if len(tmp[size]) == size {
			groups = append(groups, tmp[size])
			delete(tmp, size)
		}

	}

	return groups
}

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Println(groupThePeople([]int{2, 1, 3, 3, 3, 2}))
}
