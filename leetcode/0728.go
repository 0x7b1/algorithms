package main

import "fmt"

func isSelfDividing(n int) bool {
	for m := n; m > 0; m /= 10 {
		if !(m % 10 > 0) || n % (m % 10) > 0 {
			return false
		}
	}

	return true
}

func selfDividingNumbers(left int, right int) []int {
	var nrs []int

	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			nrs = append(nrs, i)
		}
	}

	return nrs
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
