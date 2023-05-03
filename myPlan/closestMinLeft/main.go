package main

import (
	lifo "algors/myPlan/stack"
	"fmt"
)

func main() {
	array := []int{7, 2, 4, 5, 3, 2, 5, 1, 5, 4}
	result := make([]int, len(array))
	stack := lifo.NewStack()

	for i := len(array) - 1; i >= 0; i-- {
		item := array[i]
		stackTop, _ := stack.Top()

		for !stack.IsEmpty() && stackTop[0] > item {
			result[stackTop[1]] = item
			stack.Pop()
			stackTop, _ = stack.Top()
		}
		stack.Add(item, i)
	}

	fmt.Println(array)
	fmt.Println(result)
}
