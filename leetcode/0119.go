package main

import "fmt"

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		if i == 0 || i == rowIndex {
			res[i] = 1
		} else {
			sum := 1
			for j := 1; j <= i; j++ {
				sum = sum * (rowIndex - j + 1) / j
			}

			res[i] = sum
		}
	}

	return res
}

func main() {
	fmt.Println(getRow(4))
	fmt.Println(getRow(3))
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
}
