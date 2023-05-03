package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// сжатие последовательностей [1,4,5,2,3,9,8,11,0] => "0-5,8-9,11"
func main() {
	input := []int{1, 4, 5, 2, 3, 9, 8, 11, 0}
	//input := []int{1}
	res := compress(input)
	fmt.Println(res)
}

func compress(arr []int) string {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	first := arr[0]
	var result []string
	counter := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1]+1 == arr[i] {
			counter++
		} else {
			if counter == 0 {
				result = append(result, strconv.Itoa(first))
			} else {
				result = append(result, strconv.Itoa(first)+"-"+strconv.Itoa(first+counter))
			}
			first = arr[i]
			counter = 0
		}
	}

	if counter == 0 {
		result = append(result, strconv.Itoa(first))
	} else {
		result = append(result, strconv.Itoa(first)+"-"+strconv.Itoa(first+counter))
	}

	return strings.Join(result, ",")
}
