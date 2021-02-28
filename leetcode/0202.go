package main

import "fmt"

func isHappy(n int) bool {
	if n == 0 {
		return false
	}

	sum := 0
	num := n
	memo := map[int]int{}

	for {
		for num != 0 {
			sum += (num % 10) * (num % 10)
			num /= 10
		}

		if _, ok := memo[sum]; !ok {
			if sum == 1 {
				return true
			}

			memo[sum] = sum
			num = sum
			sum = 0
		} else {
			return false
		}
	}
}

func main() {
	fmt.Println(isHappy(19))
}
