package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var limits, clicks []string
	dict := make(map[int]int)

	sc.Scan()
	sc.Scan()
	limits = strings.Fields(sc.Text())
	//
	sc.Scan()
	sc.Scan()
	clicks = strings.Fields(sc.Text())

	for _, button := range clicks {
		button, _ := strconv.Atoi(button)
		if _, ok := dict[button]; !ok {
			dict[button] = 0
		}
		dict[button]++
	}

	for i, limit := range limits {
		button := i + 1
		limit, _ := strconv.Atoi(limit)

		if limit < dict[button] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
