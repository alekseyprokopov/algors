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
	nk := toInt(sc)
	n, k := nk[0], nk[1]
	arr := toInt(sc)

	var counter = make([]int, k+1)

	for _, item := range arr {
		counter[item]++
	}

	left := 0
	right := n - 1

	for counter[arr[right]] != 1 {
		counter[arr[right]]--
		right--
	}

	for counter[arr[left]] != 1 {
		counter[arr[left]]--
		left++
	}

	fmt.Println(left, right)
	fmt.Println(arr[left], arr[right])

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
