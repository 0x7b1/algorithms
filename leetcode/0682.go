package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	stack := make([]int, 0)

	for _, op := range ops {
		if op == "+" {
			a, b := stack[len(stack) - 1], stack[len(stack) - 2]
			stack = append(stack, a + b)
		} else if op == "D" {
			a := stack[len(stack) - 1]
			stack = append(stack, a * 2)
		} else if op == "C" {
			stack = stack[:len(stack) - 1]
		} else {
			num, _ := strconv.Atoi(op)
			stack = append(stack, num)
		}
	}

	sum := 0
	for _, num := range stack {
		sum += num
	}

	return sum
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
}
