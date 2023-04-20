package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	list := strings.Split(scan.Text(), " ")
	intList := make([]int, len(list))
	for i, item := range list {
		intItem, _ := strconv.Atoi(item)
		intList[i] = intItem
	}
	counter := 0
	for i := 1; i < len(intList)-1; i++ {
		if intList[i] > intList[i-1] && intList[i] > intList[i+1] {
			counter++
		}
	}
	fmt.Println(counter)
}
