package main

import (
	"fmt"
	"sort"
)

func sumZero(n int) []int {
	var res []int

	if n%2 == 0 {
		for i := 1; i <= n/2; i++ {
			res = append(res, i, -i)
		}
	} else {
		res = append(res, 0)
		for i := 1; i <= (n-1)/2; i++ {
			res = append(res, i, -i)
		}
	}

	sort.Ints(res)

	return res
}

func main() {
	fmt.Println(sumZero(5))
	fmt.Println(sumZero(3))
}
