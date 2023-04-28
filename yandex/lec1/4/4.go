package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Scan()
	number := sc.Text()
	var checks []string
	replacer := strings.NewReplacer("(", "", ")", "", "-", "", "+7", "")

	for sc.Scan() {
		checks = append(checks, sc.Text())
	}
	number = fixNumber(number, replacer)
	for _, check := range checks {
		if number == fixNumber(check, replacer) {
			fmt.Println("YES")

		} else {
			fmt.Println("NO")

		}
	}
}

func fixNumber(number string, replacer *strings.Replacer) string {
	result := strings.TrimPrefix(number, "8")
	result = replacer.Replace(result)

	if len(result) < 10 {
		result = "495" + result
	}
	return result
}
