package main

import (
	lifo "algors/myPlan/stack"
	"fmt"
)

func main() {
	stack := lifo.NewStack()
	stack.Add(2)
	stack.Add(3)
	stack.Add(4)
	stack.Add(5)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}
