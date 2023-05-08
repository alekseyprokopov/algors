package main

import "fmt"

func main() {
	start := "VII"

	fmt.Println(romanToInt(start))

}

func romanToInt(s string) int {
	dict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var result, last int

	for i := len(s) - 1; i >= 0; i-- {
		value := dict[string(s[i])]

		if value >= last {
			result += value
		} else {
			result -= value
		}
		last = value

	}

	return result
}
