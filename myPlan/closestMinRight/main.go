package main

import (
	lifo "algors/myPlan/stack"
	"fmt"
)

func main() {
	array := []int{7, 2, 4, 5, 3, 2, 5, 1, 5, 4}
	result := make([]int, len(array))
	stack := lifo.NewStack()

	for i, item := range array {
		value, _ := stack.Top()
		for !stack.IsEmpty() && item < value[0] {
			result[value[1]] = i
			stack.Pop()
			value, _ = stack.Top()
		}
		stack.Add(item, i)
	}

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		result[value[1]] = len(array)
	}

	fmt.Println(array)
	fmt.Println(result)
}
