package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	occ := make(map[int]int)
	for _, n := range arr {
		occ[n]++
	}

	back := make(map[int]int)
	for k, v := range occ {
		back[v] = k
	}

	return len(occ) == len(back)
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 2, 1, 1, 3}))
}
