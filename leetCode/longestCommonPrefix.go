package main

import (
	"fmt"
	"strings"
)

func main() {
	start := []string{"flower"}

	fmt.Println(longestCommonPrefix(start))
}
func longestCommonPrefix(str []string) string {
	first := strings.Split(str[0], "")
	min := len(first)
	for _, word := range str {
		counter := 0
		for index, char := range first {
			if index > len(word)-1 || char != string(word[index]) {
				break
			}
			counter++
		}
		if min > counter {
			min = counter
		}
	}

	return strings.Join(first[:min], "")

}
