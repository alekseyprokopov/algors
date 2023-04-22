package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	list := strings.Split(string(data), " ")
	fmt.Println(list)
	//dict := make(map[string]int)
	//var res []string
	//for _, word := range list {
	//	res = append(res, strconv.Itoa(dict[word]))
	//	dict[word]++
	//}
	//fmt.Println(strings.Join(res, " "))

}
