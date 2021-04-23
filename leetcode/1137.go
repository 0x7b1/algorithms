package main

import "fmt"

func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 1
	d := a + b + c

	for n > 2 {
		d = a + b + c
		a, b, c = b, c, d
		n--
	}
	return c
}

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}
