package main

import "fmt"

func hammingDistance(x int, y int) int {
	diff := x ^ y
	count := 0

	for diff > 0 {
		count += diff & 1
		diff >>= 1
	}

	return count
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
