package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	var n, k int
//	fmt.Scanf("%d %d", &n, &k)
//	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
//	cars := make([]int, n)
//
//	for i, item := range strings.Fields(text) {
//		car, _ := strconv.Atoi(item)
//		cars[i] = car
//	}
//	fmt.Println(cars)
//	fmt.Println(getPrefix(cars))
//
//}
//
//func getPrefix(array []int) *[]int {
//	prefix := make([]int, len(array)+1)
//
//	for i, item := range array {
//		prefix[i+1] = prefix[i] + item
//	}
//	return &prefix
//
//}
