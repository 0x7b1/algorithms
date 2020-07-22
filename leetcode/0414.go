package main

import (
	"fmt"
	"math"
)

func thirdMax(nums []int) int {
	inf := -math.MaxInt64
	m1, m2, m3 := inf, inf, inf

	for _, num := range nums {
		if num > m1 {
			m3, m2, m1 = m2, m1, num
		} else if num > m2 && num != m1 {
			m3, m2 = m2, num
		} else if num > m3 && num != m2 && num != m1 {
			m3 = num
		}
	}

	if m3 == inf {
		return m1
	}

	return m3
}

func main() {
	fmt.Println(thirdMax([]int{1, 2}))
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
	fmt.Println(thirdMax([]int{1, 2, 3, 4, 5}))
}
