package main

import (
	"fmt"
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	time := 0
	for i := 0; i < len(points) - 1; i++ {
		num := math.Max(math.Abs(float64(points[i+1][0]-points[i][0])), math.Abs(float64(points[i+1][1]-points[i][1])))
		time += int(num)
	}

	return time
}

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
}
