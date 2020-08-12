package main

import "fmt"

func findComplement(num int) int {
	digits := 0
	tmp := num
	for tmp > 0 {
		tmp >>= 1
		digits = (digits << 1) | 1
	}

	return num ^ digits
}

func main() {
	fmt.Println(findComplement(5))
}
