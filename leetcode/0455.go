package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var i int

	for j := 0; j < len(s) && i < len(g); j++ {
		if g[i] <= s[j] {
			i++
		}
	}

	return i
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println(findContentChildren([]int{1, 1}, []int{2}))
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{3}))
}
