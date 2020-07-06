package main

import "fmt"

func numberOfSteps(num int) int {
	steps := 0
	for num != 0 {
		if num&1 == 0 {
			steps += 1
		} else {
			steps += 2
		}
		num >>= 1
	}

	if steps == 0 {
		steps++
	}

	return steps - 1
}

func main() {
	fmt.Println(numberOfSteps(14))
}
