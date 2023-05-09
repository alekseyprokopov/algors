package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
}

func lengthOfLastWord(s string) int {
	data := strings.Fields(s)
	return len(data[len(data)-1])
}
