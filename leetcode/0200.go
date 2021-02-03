package main

import "fmt"

func dfs(grid [][]byte, j, i, width, height int) {
	if j < 0 || j >= height || i < 0 || i >= width || grid[j][i] != '1' {
		return
	}

	grid[j][i] = 0
	dfs(grid, j, i+1, width, height)
	dfs(grid, j, i-1, width, height)
	dfs(grid, j+1, i, width, height)
	dfs(grid, j-1, i, width, height)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0

	width, height := len(grid[0]), len(grid)

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if grid[j][i] == '1' {
				dfs(grid, j, i, width, height)
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}))

	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}))
}
