package main

import (
	"fmt"
	"math/bits"
)

func isPrime(x int) bool {
	return x == 2 || x == 3 || x == 5 || x == 7 || x == 11 ||
		x == 13 || x == 17 || x == 19
}

func countPrimeSetBits(L int, R int) int {
	res := 0
	for i := L; i <= R; i++ {
		if isPrime(bits.OnesCount(uint(i))) {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(10, 15))
}
