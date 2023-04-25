package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	first, second := make([]int, n), make([]int, k)

	firstText, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	secondText, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	for i, item := range strings.Fields(firstText) {
		itemInt, _ := strconv.Atoi(item)
		first[i] = itemInt
	}
	for i, item := range strings.Fields(secondText) {
		itemInt, _ := strconv.Atoi(item)
		second[i] = itemInt
	}

	for _, item := range second {
		ans := binSearch(first, item)
		if ans != -1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func binSearch(array []int, check int) int {
	l := 0
	r := len(array) - 1
	for l < r {
		m := (r + l) / 2
		if array[m] >= check {
			r = m
		} else {
			l = m + 1
		}

	}
	if array[l] == check {
		return l
	}
	return -1
}
