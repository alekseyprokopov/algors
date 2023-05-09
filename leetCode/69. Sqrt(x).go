package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(10))
}

func mySqrt(x int) int {
	l := 0
	r := x

	for l < r {
		mid := (l + r + 1) / 2

		if mid*mid <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
