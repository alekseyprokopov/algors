package main

import (
	"fmt"
	"strconv"
	"strings"
)

// RLE (сжатие)
func main() {
	//input := "AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBB"
	input := "A"
	fmt.Println(RLE(input))
}

func RLE(text string) string {
	chars := strings.Split(text, "")
	var result []string
	counter := 1
	last := chars[0]
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			counter++
		} else {
			digit := strconv.Itoa(counter)
			result = append(result, chars[i-1], digit)
			counter = 1
		}
		last = chars[i]
	}
	result = append(result, last, strconv.Itoa(counter))
	return strings.Join(result, "")

}
