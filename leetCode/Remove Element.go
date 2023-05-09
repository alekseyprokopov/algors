package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	counter := 0
	for i, num := range nums {
		if num != val {
			nums[counter] = nums[i]
			counter++
		}
	}
	return counter
}
