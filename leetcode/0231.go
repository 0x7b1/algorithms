package main

import "fmt"

func isPowerOfTwoNaive(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}

	return true
}

func isPowerOfTwo(n int) bool {
	// return n > 0 && ((n & (n - 1)) == 0)
	return n > 0 && (1073741824%n == 0)
}

func main() {
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(-16))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(5))
	fmt.Println(isPowerOfTwo(32))
	fmt.Println(isPowerOfTwo(30))
}
