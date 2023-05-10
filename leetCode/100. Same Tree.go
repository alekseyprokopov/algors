package main

import "fmt"

func main() {
	test := []int{}
	test = append(test, nil...)
	fmt.Println(test)
}
