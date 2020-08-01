package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	counter := 0
	neighbors := [][]int{
		{-1, 0}, // left
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
	}
	for i, row := range grid {
		for j, e := range row {
			if e == 1 {
				if i == 0 || i == len(grid)-1 || j == 0 || j == len(row)-1 {
					counter++
				}

				for _, side := range neighbors {
					ii, jj := i+side[0], j+side[1]

					if ii < 0 || ii > len(grid)-1 || jj < 0 || jj > len(row)-1 {
						continue
					}

					if grid[ii][jj] == 0 {
						counter++
					}
				}
			}
		}
	}

	return counter
}

func main() {
	fmt.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
}
