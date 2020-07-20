package main

import "fmt"

/*
 * Complete the 'dynamicArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */
func dynamicArray(n int32, queries [][]int32) []int32 {
	array := make([][]int32, n)
	var rsp []int32
	var lastAnswer int32

	for _, query := range queries {
		queryNr, x, y := query[0], query[1], query[2]

		idx := (x ^ lastAnswer) % n
		if queryNr == 1 {
			array[idx] = append(array[idx], y)
		} else if queryNr == 2 {
			valIdx := y % int32(len(array[idx]))
			lastAnswer = array[idx][valIdx]
			rsp = append(rsp, lastAnswer)
		}
	}

	return rsp
}

func main() {
	fmt.Println(dynamicArray(2, [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}))
}
