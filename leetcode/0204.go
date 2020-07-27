package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func countPrimes(n int) int {
	counter := 0
	if n <= 2 {
		return 0
	}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(countPrimes(10))
}
