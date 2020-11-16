package main

import (
	"fmt"
	"github.com/pkg/math"
)

func minDistanceRecursive(word1, word2 string) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == - 1 {
			return j + 1
		}

		if j == -1 {
			return i + 1
		}

		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		} else {
			insertOp := dp(i, j-1) + 1
			deleteOp := dp(i-1, j) + 1
			replaceOp := dp(i-1, j-1) + 1

			return math.MinN(insertOp, deleteOp, replaceOp)
		}
	}

	return dp(len(word1)-1, len(word2)-1)
}

func minDistanceMemo(word1, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Base case
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	// From the bottom up
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				insertOp := dp[i][j-1] + 1
				deleteOp := dp[i-1][j] + 1
				replaceOp := dp[i-1][j-1] + 1

				dp[i][j] = math.MinN(insertOp, deleteOp, replaceOp)
			}
		}
	}

	// Store the least editing distance from word1 to word2
	return dp[m][n]
}

func main() {
	word1 := "horse"
	word2 := "rse"

	dst1 := minDistanceRecursive(word1, word2)
	fmt.Println(dst1)

	dst2 := minDistanceMemo(word1, word2)
	fmt.Println(dst2)
}
