package main

import "fmt"

func isBadVersion(n int) bool {
	return false
}

func firstBadVersion(n int) int {
	lo, hi := 1, n

	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func main() {
	fmt.Println(firstBadVersion(5))
}
