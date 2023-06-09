package main

import "fmt"

func main() {
	//digits := []int{1, 2, 3}
	digits := []int{9, 8}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0

	}
	digits = append([]int{1}, digits...)
	return digits

}
