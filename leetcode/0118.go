package main

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		row := make([]int, 0)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else if i > 1 {
				row = append(row, res[i-1][j-1]+res[i-1][j])
			}
		}

		res = append(res, row)
	}

	return res
}

func main() {
	fmt.Println(generate(5))
}
