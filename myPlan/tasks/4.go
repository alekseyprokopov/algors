package main

import "fmt"

//intersection

func main() {
	arr1 := []int{1, 2, 3, 2, 0}
	arr2 := []int{5, 1, 2, 7, 3, 2}

	dict1 := map[int]int{}
	var result []int
	for _, item := range arr1 {
		dict1[item]++
	}

	for _, item := range arr2 {
		if count, ok := dict1[item]; ok {
			result = append(result, item)
			count--
		}
	}

	fmt.Println(result)

}
