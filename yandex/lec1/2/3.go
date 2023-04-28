package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var number string
	fmt.Scan(&number)

	res := strings.TrimFunc(number, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	fmt.Println(res)
}
