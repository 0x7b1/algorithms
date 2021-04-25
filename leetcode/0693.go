package main

import "fmt"

func hasAlternatingBits2(n int) bool {
	prev := n & 1
	n >>= 1

	for n > 0 {
		if prev == n & 1 {
			return false
		}

		prev = n & 1

		n >>= 1
	}

	return true
}

func hasAlternatingBits(n int) bool {
	v := 1 + n + n >> 1
	return v & (v - 1) == 0
}

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(10))
	fmt.Println(hasAlternatingBits(3))
}
