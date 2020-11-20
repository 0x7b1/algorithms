package main

import (
	"fmt"
	"github.com/pkg/math"
)

func fourKeysRecursive(n int) int {
	var dp func(a, aNum, copy int) int

	dp = func(n, aNum, copy int) int {
		if n <= 0 {
			return aNum
		}

		return math.MaxN(
			dp(n-1, aNum+1, copy),    // A
			dp(n-1, aNum+copy, copy), // Ctrl-V
			dp(n-2, aNum, aNum),      // Ctrl-C + Ctrl+V
		)
	}

	return dp(n, 0, 0)
}

func fourKeysDP(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := 2; j < i; j++ {
			dp[i] = math.MaxN(dp[i], dp[j-2]*(i-j+1))
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(fourKeysRecursive(3), fourKeysDP(3))
	fmt.Println(fourKeysRecursive(7), fourKeysDP(7))
}
