package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	set := make(map[int]int)

	for _, n := range nums1 {
		set[n]++
	}

	res := make([]int, 0)
	for _, n := range nums2 {
		if _, ok := set[n]; ok {
			res = append(res, n)
			set[n]--
			if set[n] == 0 {
				delete(set, n)
			}
		}
	}

	return res
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
