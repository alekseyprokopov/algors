package main

import "fmt"

func main() {
	first := []int{1, 3, 5, 8, 9, 9}
	second := []int{10, 2, 4, 8, 1, 6}

	for _, item := range second {
		ans := lBinSearch(first, item)
		if ans != len(first) && (ans == 0 || (Abs(first[ans]-item) < Abs(first[ans-1]-item))) {
			fmt.Println(first[ans])
		} else {
			fmt.Println(first[ans-1])
		}
	}
}

func lBinSearch(array []int, value int) int {
	l := 0
	r := len(array)
	for l < r {
		m := (l + r) / 2
		if array[m] > value {
			r = m
		} else if array[m] < value {
			l = m + 1
		} else {
			return m
		}
	}
	return l
}
func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
