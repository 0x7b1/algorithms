package main

import "fmt"

func arrayManipulation(n int32, queries [][]int32) int64 {
	var max int64
	agg := make([]int64, n+1)
	for _, query := range queries {
		a, b, k := query[0], query[1], query[2]
		agg[a] += int64(k)
		if b+1 <= n {
			agg[b+1] -= int64(k)
		}
	}

	var x int64
	for i := 1; i <= int(n); i++ {
		x += agg[i]
		if max < x {
			max = x
		}

	}
	return max
}

func main() {
	fmt.Println(arrayManipulation(5, [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}}))
}
