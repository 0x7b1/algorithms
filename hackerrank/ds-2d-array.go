package main

import "fmt"

func hourglassSum(arr [][]int32) int32 {
	indx := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 0}, {1, -1}, {1, 0}, {1, 1}}
	var sum, max int32

	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			for _, idx := range indx {
				sum += arr[i+idx[0]][j+idx[1]]
			}

			if i == 1 && j == 1 {
				max = sum
			} else if sum > max {
				max = sum
			}

			sum = 0
		}
	}

	return max
}

func main() {
	fmt.Println(hourglassSum([][]int32{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}))

	fmt.Println(hourglassSum([][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}))
}
