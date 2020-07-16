package main

import (
	"fmt"
	"sort"
	"strings"
)

func minDeletionSize(A []string) int {
	result := 0
	for c := 0; c < len(A[0]); c++ {
		for r := 1; r < len(A); r++ {
			if A[r - 1][c] > A[r][c] {
				result++
				break
			}
		}
	}

	return result
}

func minDeletionSizeDumb(A []string) int {
	n := len(A[0])
	var seq []string
	tmp := ""
	var chars []string

	for i := 0; i < n; i++ {
		for j := 0; j < len(A); j++ {
			tmp += string(A[j][i])
		}

		chars = strings.Split(tmp, "")
		if !sort.IsSorted(sort.StringSlice(chars)) {
			seq = append(seq, tmp)
		}

		tmp = ""
	}

	return len(seq)
}

// - 0 c, d, g
// x 1 b, a, h
// - 2 a, f, i

// x 0 z w t
// x 1 y v s
// x 2 x u r

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
	fmt.Println(minDeletionSize([]string{"a", "b"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
