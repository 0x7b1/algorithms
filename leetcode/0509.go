package main

import "fmt"

func fib(N int) int {
	a, b := 1, 1

	if N == 0 {
		return 0
	}

	for N > 2 {
		a, b = b, a + b
		N--
	}

	return b
}

func main() {
	fmt.Println(fib(3))
	fmt.Println(fib(4))
}
