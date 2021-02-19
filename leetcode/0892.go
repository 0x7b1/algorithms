package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func surfaceArea(grid [][]int) int {
	res, n := 0, len(grid)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				res += grid[i][j]*4 + 2
			}

			if i > 0 {
				res -= min(grid[i][j], grid[i-1][j]) * 2
			}

			if j > 0 {
				res -= min(grid[i][j], grid[i][j-1]) * 2
			}
		}
	}

	return res
}

func main() {
	fmt.Println(surfaceArea([][]int{{2}}))
	fmt.Println(surfaceArea([][]int{{1, 2}, {3, 4}}))
}
