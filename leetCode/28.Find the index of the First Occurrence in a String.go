package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strStr("leetcode", "code"))
	fmt.Println(strStr2("leetcode", "code"))
}

func strStr(haystack string, needle string) int {
	text := strings.Split(haystack, "")
	target := strings.Split(needle, "")

	for i, _ := range text {
		left := i
		right := 0
		counter := 0
		for right < len(target) && left < len(text) && text[left] == target[right] {
			counter++
			left++
			right++
			if counter == len(target) {
				return i
			}

		}

	}
	return -1
}

func strStr2(haystack string, needle string) int {

	for i := 0; i < len(haystack); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
