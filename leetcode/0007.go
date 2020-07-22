package main

import "fmt"

func reverse(x int) int {
	if x < 0 {
		return -1 * reverse(-x)
	}

	num := 0
	for x > 0 {
		num = num * 10 + x % 10
		x /= 10
	}

	if num <= 0x7fffffff {
		return num
 	} else {
 		return 0
	}
}

func main() {
	fmt.Println(reverse(-123))
	fmt.Println(reverse(123))
	fmt.Println(reverse(120))
}
