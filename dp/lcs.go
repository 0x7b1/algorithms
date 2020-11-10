package main

import "fmt"

// This is a bruteforce solution
func longestCommonSubsequenceRecursion(str1, str2 string) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 || j == - 1 {
			return 0
		}

		if str1[i] == str2[j] {
			// found a character that belongs to lcs, keep finding
			return dp(i-1, j-1) + 1
		} else {
			left := dp(i-1, j)
			up := dp(i, j-1)

			if left > up {
				return left
			}

			return up
		}
	}

	return dp(len(str1)-1, len(str2)-1)
}

func longestCommonSubsequenceMemo(str1, str2 string) int {
	m, n := len(str1), len(str2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				left := dp[i-1][j]
				up := dp[i][j-1]

				if left > up {
					dp[i][j] = left
				} else {
					dp[i][j] = up
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	str1 := "babcde"
	str2 := "acze"

	longestStr1 := longestCommonSubsequenceRecursion(str1, str2)
	fmt.Println(longestStr1)

	longestStr2 := longestCommonSubsequenceMemo(str1, str2)
	fmt.Println(longestStr2)
}
