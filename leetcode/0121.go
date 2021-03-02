package main

import "fmt"

func maxProfitBruteForce(prices []int) int {
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	min, max := prices[0], 0

	for _, price := range prices {
		if price - min > max {
			max = price - min
		} else if price < min {
			min = price
		}
	}

	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
