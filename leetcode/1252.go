package main

import "fmt"

func oddCells2(n int, m int, indices [][]int) int {
	cells := make([]int, n*m)

	for _, pair := range indices {
		// Incrementing rows
		for i := 0; i < m; i++ {
			cells[m*pair[0]+i]++
		}

		// Incrementing columns
		for j := 0; j < n; j++ {
			cells[m*j+pair[1]]++
		}
	}

	counter := 0

	for _, cell := range cells {
		if cell%2 != 0 {
			counter++
		}
	}

	return counter
}

func oddCells(n int, m int, indices [][]int) int {
	row, col := make([]int, n), make([]int, m)
	for _, idx := range indices {
		row[idx[0]] ^= 1
		col[idx[1]] ^= 1
	}

	row_sum, col_sum := 0, 0
	for _, r := range row {
		row_sum += r
	}

	for _, c := range col {
		col_sum += c
	}

	return row_sum * m + col_sum * n - 2 * row_sum * col_sum
}

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}}))
}
