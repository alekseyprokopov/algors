package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	tar := 7
	fmt.Println(searchInsert(nums, tar))
}

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r) / 2

		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if nums[l] < target {
		return len(nums)
	} else {
		return l
	}
}
