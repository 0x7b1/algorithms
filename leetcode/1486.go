package main

import "fmt"

func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + 2 * i
	}

	return res
}

func main() {
	fmt.Println(xorOperation(10, 20))
}
