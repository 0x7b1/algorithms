package main

import "fmt"

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}

	return n / 5 + trailingZeroes(n / 5)
}

func main() {
	fmt.Println(trailingZeroes(3))
	fmt.Println(trailingZeroes(5))
	fmt.Println(trailingZeroes(0))
}
