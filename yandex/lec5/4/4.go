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
	sc.Scan()
	nk := toInt(strings.Fields(sc.Text()))
	_, k := nk[0], nk[1]

	sc.Scan()
	array := toInt(strings.Fields(sc.Text()))

	res := find(array, k)
	fmt.Println(res)
}

func find(arr []int, r int) int {
	counter := 0
	j := 1
	for i := range arr {
		for j < len(arr) && arr[j]-arr[i] <= r {
			j++
		}
		counter += len(arr) - j
	}
	return counter
}

func toInt(array []string) []int {
	var res []int
	for _, item := range array {
		intItem, _ := strconv.Atoi(item)
		res = append(res, intItem)
	}
	return res
}

//func find(arr []int, r int) int {
//	counter := 0
//	i, j := 0, 1
//	for i < len(arr) && j < len(arr) {
//		if arr[j]-arr[i] > r {
//			counter += len(arr) - j
//			i++
//		} else {
//			j++
//		}
//	}
//	return counter
//}
