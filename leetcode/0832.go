package main

import "fmt"

func flipAndInvertImage2(A [][]int) [][]int {
	cols := len(A[0])
	for i := 0; i < len(A); i++ {
		for j := 0; j < cols/2; j++ {
			A[i][j], A[i][cols-1-j] = A[i][cols-1-j], A[i][j]
		}
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] ^=1
		}
	}

	return A
}

func flipAndInvertImage(A [][]int) [][]int {
	n := len(A)
	for _, row := range A {
		for i := 0; i * 2 < n; i++ {
			if row[i] == row[n - i - 1] {
				row[n - i - 1] ^= 1
				row[i] = row[n - i - 1]
			}
		}
	}

	return A
}

func main() {
	fmt.Println(flipAndInvertImage([][]int{{0, 1}, {0, 1}}))
	fmt.Println(flipAndInvertImage([][]int{{1, 0, 0}, {0, 1, 1}}))
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}
