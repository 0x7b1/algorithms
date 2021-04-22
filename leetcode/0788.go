package main

import "fmt"

func rotatedDigits(N int) int {
	dp := make([]int, N + 1)
	count := 0

	for i := 0; i <= N; i++ {
		if i < 10 {
			if i == 0 || i == 1 || i == 8 {
				dp[i] = 1
			} else if i == 2 || i == 5 || i == 6 || i == 9 {
				dp[i] = 2
				count++
			}
		} else {
			a := dp[i / 10]
			b := dp[i % 10]
			if a == 1 && b == 1 {
				dp[i] = 1
			} else if a >= 1 && b >= 1 {
				dp[i] = 2
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(rotatedDigits(10))
}
