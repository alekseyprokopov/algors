package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	toInt(sc)

	first := toInt(sc)
	second := toInt(sc)

	for _, item := range second {
		fmt.Println(findClosest(first, item))
	}

}

func toInt(scanner *bufio.Scanner) []int {
	scanner.Scan()
	array := strings.Split(scanner.Text(), " ")
	var res []int
	for _, item := range array {
		intItem, _ := strconv.Atoi(item)
		res = append(res, intItem)
	}
	return res
}

func lBinSearch(arr []int, value int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		m := (l + r) / 2
		if arr[m] >= value {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func rBinSearch(arr []int, value int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		mid := (l + r + 1) / 2
		if arr[mid] <= value {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func findClosest(arr []int, value int) int {
	lSearch := lBinSearch(arr, value)
	rSearch := rBinSearch(arr, value)
	if abs(arr[rSearch]-value) > abs(arr[lSearch]-value) {
		return arr[lSearch]
	}
	return arr[rSearch]

}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
