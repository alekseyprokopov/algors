package main

import "fmt"

func main() {
	input := []int{3, 3}
	target := 6

	fmt.Println(twoSum(input, target))
}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	var result []int
	for i, item := range nums {
		if _, ok := dict[item]; ok {
			result = append(result, dict[item], i)
		}
		dict[target-item] = i
	}
	return result
}
