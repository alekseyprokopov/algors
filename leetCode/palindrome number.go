package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(0))
}

func isPalindrome(x int) bool {
	data := []rune(strconv.Itoa(x))
	for i, j := 0, len(data)-1; i < len(data)/2; i, j = i+1, j-1 {
		if data[i] != data[j] {
			return false
		}
	}
	return true
}
