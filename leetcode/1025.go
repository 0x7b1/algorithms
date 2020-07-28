package main

import "fmt"

func divisorGame(N int) bool {
	count := 0
	for N > 1 {
		flag := true
		for i := 1; i < N; i++ {
			if N % 1 == 0 {
				flag = false
				N -= i
				break
			}
		}

		if flag {
			return false
		}

		count++
	}

	return count % 2 != 0
}

func main() {
	fmt.Println(divisorGame(2))
	fmt.Println(divisorGame(3))
}
