package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	m := len(matrix)
	n := len(matrix[0])
	border := [][]int{{0, 0}}

	// Collect border positions
	for j := 1; j < n; j++ {
		border = append(border, []int{0, j})
	}

	for i := 1; i < m; i++ {
		border = append(border, []int{i, 0})
	}

	// Search for valid diagonals
	for _, pos := range border {
		i, j := pos[0], pos[1]
		nr := matrix[i][j]

		for {
			if j == n || i == m {
				break
			}


			if matrix[i][j] != nr {
				return false
			}

			i++
			j++
		}
	}

	return true
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 1, 2, 3},
		{9, 5, 1, 2},
	}
	fmt.Println(isToeplitzMatrix(matrix))
}
