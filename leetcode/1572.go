package main

import "fmt"

func diagonalSum(mat [][]int) int {
	n := len(mat)
	var sum int

	for i := 0; i < n; i++ {
		sum += mat[i][i]
		sum += mat[i][n - i - 1]
	}

	if n % 2 != 0 {
		mid := (n - 1) / 2
		sum -= mat[mid][mid]
	}

	return sum
}

func main() {
	fmt.Println(diagonalSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(diagonalSum([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}))
}
