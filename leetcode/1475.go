package main

import "fmt"

func finalPrices(prices []int) []int {
	stack := []int{}

	for i := 0; i < len(prices); i++ {
		for len(stack) > 0 &&  prices[stack[len(stack) - 1]] >= prices[i] {
			prices[stack[len(stack) - 1]] -= prices[i]
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, i)
	}

	return prices
}

func main() {
	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}
