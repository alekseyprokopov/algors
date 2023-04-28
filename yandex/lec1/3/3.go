package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d\n%d\n%d\n", &a, &b, &c)
	if c < 0 {
		fmt.Println("NO SOLUTION")
	} else if a == 0 {
		fmt.Println("MANY SOLUTIONS")
	} else {
		res := (c*c - b) / a
		fmt.Print(res)
	}
}
