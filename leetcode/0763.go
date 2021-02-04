package main

import "fmt"

func partitionLabels(S string) []int {
	res := make([]int, 0)
	lastIndex := make(map[rune]int)

	for i, c := range S {
		lastIndex[c] = i
	}

	var j, anchor int
	for i, c := range S {
		last := lastIndex[c]
		if last > j {
			j = last
		}

		if i == j {
			res = append(res, i-anchor+1)
			anchor = i + 1
		}
	}

	return res
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
