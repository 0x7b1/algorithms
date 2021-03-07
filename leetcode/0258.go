package main

import "fmt"

// Time: O(n)
func addDigitsWithLoop(num int) int {
	for num > 9 {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}

		num = sum
	}

	return num
}

// Time: O(1)
func addDigits(num int) int {
	if num == 0 {
		return 0
	}

	return 1 + (num - 1) % 9
}

func main() {
	fmt.Println(addDigits(38))
}
