package main

import (
	"fmt"
	"sort"
)

type Intervals [][]int

func (a Intervals) Len() int {
	return len(a)
}

func (a Intervals) Less(i, j int) bool {
	return a[i][1] < a[j][1]
}

func (a Intervals) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func intervalScheduling(vals [][]int) int {
	if len(vals) == 0 {
		return 0
	}

	// Sort in ascending order by the ending time
	sort.Sort(Intervals(vals))

	// At least one interval without intersection
	count := 1

	// After sorting the first interval is x
	x_end := vals[0][1]

	for _, interval := range vals {
		start := interval[0]
		if start >= x_end {
			count++
			x_end = interval[1]
		}
	}

	return count
}

func main() {
	vals := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}

	res := intervalScheduling(vals)
	fmt.Println(res)
}
