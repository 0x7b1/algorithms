// ---------------------------------------
// Find the missing number from a sequence
// ---------------------------------------
// Given an array of length n containing numbers from 0,...,n
// find out the missing number from the sequence
//
// Example:
// Input: [3, 0, 1]
// Output: 2

package main

import (
	"fmt"
	"sort"
)

// Time: O(nlogn), Space: O(n)
func findMissingNumberViaSorting(arr []int) int {
	sort.Ints(arr)
	count := 0
	for _, n := range arr {
		if count != n {
			return count
		}

		count++
	}

	return -1
}

// Time: O(n), Space: O(n)
func findMissingNumberViaHashset(arr []int) int {
	hash := make(map[int]bool)
	for _, n := range arr {
		hash[n] = true
	}

	for i := 0; i <= len(arr); i++ {
		if _, ok := hash[i]; !ok {
			return i
		}
	}

	return -1
}

// Time: O(n), Space: O(1)
func findMissingNumberViaXOR(arr []int) int {
	n := len(arr)
	res := n

	for i := 0; i < n; i++ {
		res ^= i ^ arr[i]
	}

	return res
}

// Time: O(n), Space: O(1)
func findMissingNumberViaArithmeticProgression(arr []int) int {
	n := len(arr)
	expect := n * (n + 1) / 2
	sum := 0

	for _, n := range arr {
		sum += n
	}

	return expect - sum
}

// Time: O(n), Space: O(1)
func findMissingNumberViaDifference(arr []int) int {
	n := len(arr)
	res := n

	for i := 0; i < n; i++ {
		res += i - arr[i]
	}

	return res
}

func main() {
	//arr := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	arr := []int{0, 3, 1, 4}

	fmt.Println(findMissingNumberViaSorting(arr))
	fmt.Println(findMissingNumberViaHashset(arr))
	fmt.Println(findMissingNumberViaXOR(arr))
	fmt.Println(findMissingNumberViaArithmeticProgression(arr))
	fmt.Println(findMissingNumberViaDifference(arr))
}
