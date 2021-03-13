package main

import "fmt"

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := (lo + hi) / 2
		if guess(mid) == 1 {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

func main() {
	fmt.Println(guessNumber(10))
	fmt.Println(guessNumber(1))
	fmt.Println(guessNumber(2))
}
