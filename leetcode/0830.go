package main

import "fmt"

func largeGroupPositions2(s string) [][]int {
	res := make([][]int, 0)
	slow, fast := 0, 1

	for fast < len(s) {
		if fast == len(s)-1 && s[slow] == s[fast] {
			if fast-slow >= 2 {
				res = append(res, []int{slow, fast})
			}
			slow = fast
		} else if s[slow] != s[fast] {
			if fast-slow > 2 {
				res = append(res, []int{slow, fast - 1})
			}
			slow = fast
		}

		fast++
	}

	return res
}

func largeGroupPositions(s string) [][]int {
	res, end := [][]int{}, 0

	for end < len(s) {
		start, str := end, s[end]
		for end < len(s) && s[end] == str {
			end++
		}

		if end-start > 2 {
			res = append(res, []int{start, end - 1})
		}
	}

	return res
}

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(largeGroupPositions("abc"))
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
	fmt.Println(largeGroupPositions("aba"))
	fmt.Println(largeGroupPositions("aaa"))
	fmt.Println(largeGroupPositions("a"))
	fmt.Println(largeGroupPositions("babaaaabbb"))
}
