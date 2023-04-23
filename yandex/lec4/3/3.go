package lec5

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	list := strings.Fields(string(data))
	dict := make(map[string]int)

	var max int
	var result string
	for _, item := range list {
		dict[item]++
		if max < dict[item] {
			max = dict[item]
			result = item
		}
	}
	fmt.Println(result)
}
