package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")

	list := strings.Fields(string(data))
	dict := make(map[string]int)
	var res []string
	for i, word := range list {
		log.Println(i, word)
		res = append(res, strconv.Itoa(dict[word]))
		dict[word]++
	}
	fmt.Println(strings.Join(res, " "))

}
