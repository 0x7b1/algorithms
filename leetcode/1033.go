package main

import (
	"fmt"
	"sort"
)

func numMovesStones(a int, b int, c int) []int {
	s := []int{a, b, c}

	sort.Ints(s)

	if s[2] - s[0] == 2 {
		return []int{0, 0}
	}

	x := 1

	if min(s[1] - s[0], s[2] - s[1]) > 2 {
		x = 2
	}

	y := s[2] - s[0] - 2

	return []int{x, y}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(numMovesStones(1, 2, 5))
}
