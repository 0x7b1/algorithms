package main

import "fmt"

func isPowerOfThree1(n int) bool {
	return n > 0 && 1162261467 % n == 0
}

func isPowerOfThree(n int) bool {
	start := 1

	for start < n {
		start *= 3
	}

	return n == start
}

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45))
}
