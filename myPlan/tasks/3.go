package main

import (
	"fmt"
	"strings"
)

//есть строка s
//подстрока из уникальных символов
//максимальная длина

func main() {
	input := "asdcasraaszzsf"
	fmt.Println(longest(input))

}

func longest(input string) int {
	chars := strings.Split(input, "")
	dict := map[string]interface{}{}
	counter := 0
	max := 0
	for _, item := range chars {
		if _, ok := dict[item]; !ok {
			counter++
			dict[item] = item
		} else {
			counter = 0
			dict = map[string]interface{}{}
		}

		if counter > max {
			max = counter
		}
	}
	return max

}
