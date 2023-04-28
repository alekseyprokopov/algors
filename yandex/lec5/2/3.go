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
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	cars := make([]int, n)

	for i, item := range strings.Fields(text) {
		car, _ := strconv.Atoi(item)
		cars[i] = car
	}
	prefix := getPrefix(cars)
	counter := 0
	for key := range prefix {
		if _, ok := prefix[key+k]; ok {
			counter += 1
		}
	}

	fmt.Println(counter)

}

func getPrefix(array []int) map[int]bool {
	prefix := make(map[int]bool)
	prefix[0] = true
	sum := 0

	for _, item := range array {
		sum += item
		prefix[sum] = true
	}
	return prefix

}
