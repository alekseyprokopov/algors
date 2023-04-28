package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	list := strings.Fields(string(data))
	dict := make(map[string]int)

	var max int
	var res []string
	for _, item := range list {
		dict[item]++
		if max < dict[item] {
			max = dict[item]
		}
	}

	for key, _ := range dict {
		if dict[key] == max {
			res = append(res, key)
		}
	}
	sort.Strings(res)
	fmt.Println(res[0])
}
