package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	first, second := 1, 1

	for i := 0; i < n; i++ {
		tmp := first + second
		first = second
		second = tmp
	}
	return first
}
