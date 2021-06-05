package main

import (
	"fmt"
	"sort"
)

var freq = make(map[int]int)

type freqArray []int

func (f freqArray) Len() int {
	return len(f)
}

func (f freqArray) Less(i, j int) bool {
	if freq[f[i]] == freq[f[j]] {
		return f[i] > f[j]
	}

	return freq[f[i]] < freq[f[j]]
}

func (f freqArray) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func frequencySort(nums []int) []int {
	freq = make(map[int]int)

	for _, n := range nums {
		freq[n]++
	}

	fmt.Println(freq)

	arr := freqArray(nums)
	sort.Sort(arr)

	return arr
}

func main() {
	fmt.Println(frequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(frequencySort([]int{2, 3, 1, 3, 2}))
}
