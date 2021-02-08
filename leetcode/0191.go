package main

import "fmt"

func hammingWeight2(num uint32) int {
	sum := 0

	for num > 0 {
		if num&1 == 1 {
			sum++
		}

		num >>= 1
	}

	return sum
}

func hammingWeight(num uint32) int {
	sum := 0
	for num != 0 {
		sum++
		num &= num - 1
	}

	return sum
}

func main() {
	fmt.Println(hammingWeight(0b00000000000000000000000000001011))
	fmt.Println(hammingWeight(0b00000000000000000000000010000000))
	fmt.Println(hammingWeight(0b11111111111111111111111111111101))
}
