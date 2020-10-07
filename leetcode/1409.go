package main

import (
	"fmt"
)

func processQueries(queries []int, m int) []int {
	seq := make([]int, m)
	res := make([]int, len(queries))

	for i := 0; i < m; i++ {
		seq[i] = i + 1
	}

	for i, p := range queries {
		j := -1
		for tmp, v := range seq {
			if p == v {
				j = tmp
				break
			}
		}

		res[i] = j
		seq = append([]int{p}, append(seq[:j], seq[j+1:]...)...)
	}

	return res
}

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}
