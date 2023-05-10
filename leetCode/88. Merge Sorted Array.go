package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	m, n = m-1, n-1
	index := len(nums1) - 1

	for m >= 0 && n >= 0 {

		if nums1[m] >= nums2[n] {
			nums1[index] = nums1[m]
			m--
		} else {
			nums1[index] = nums2[n]
			n--
		}
		index--
	}

	for m >= 0 {
		nums1[index] = nums1[m]
		m--
		index--
	}

	for n >= 0 {
		nums1[index] = nums2[n]
		n--
		index--
	}

}
