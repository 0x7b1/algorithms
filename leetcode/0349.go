package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, n := range nums1 {
		set[n] = true
	}

	res := []int{}
	for _, n2 := range nums2 {
		if _, ok := set[n2]; ok {
			res = append(res, n2)
			delete(set, n2)
		}
	}

	return res
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
