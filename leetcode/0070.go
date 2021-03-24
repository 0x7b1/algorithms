package main

import "fmt"

func climbStairs(n int) int {
	a, b := 1, 1

	for n > 0 {
		a, b = b, a + b
		n--
	}

 	return a
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
}
