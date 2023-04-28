package main

import (
	"fmt"
)

func main() {
	asd := ""
	fmt.Println(findPalindrome(asd))

	//result, c := findPalindrome(asd)

	//fmt.Println(result, c)
	//fmt.Println(asd[0])
	//fmt.Println(reflect.TypeOf(asd[0]))
}

func findPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		fmt.Println("ij: ", runes[i], runes[j])
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
