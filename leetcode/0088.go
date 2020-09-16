package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; n > 0; i-- {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			m--
			nums1[i] = nums1[m]
		} else {
			n--
			nums1[i] = nums2[n]
		}
	}
}

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}

	merge(a, len(a)-len(b), b, len(b))
	fmt.Println(a)
}
