package main

import "fmt"

func isPowerOfFour(num int) bool {
	return num > 0 && num & (num - 1) == 0 && 0b1010101010101010101010101010101 & num == num
}

func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(5))
}
