package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	lo, hi := 0, len(letters)

	for lo < hi {
		mid := lo + (hi - lo) / 2
		if letters[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return letters[lo % len(letters)]
}

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))
}
