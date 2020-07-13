package main

import "fmt"

func countNegatives(grid [][]int) int {
	counter := 0
	c := len(grid[0]) - 1

	for _, row := range grid {
		for c >= 0 && row[c] < 0 {
			c--
		}
		counter += len(grid[0]) - 1 - c
	}

	return counter
}

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
}
