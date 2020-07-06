package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	num := 0

	for _, j := range J {
		for _, s := range S {
			if j == s {
				num++
			}
		}
	}

	return num
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
