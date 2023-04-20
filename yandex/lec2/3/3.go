package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	fmt.Scanf("%d\n", &n)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	list := strings.Split(scan.Text(), " ")
	intList := make([]int, len(list))
	for i, item := range list {
		intItem, _ := strconv.Atoi(item)
		intList[i] = intItem
	}
	fmt.Scanf("%d\n", &k)

	value := intList[0]
	diff := math.Abs(float64(k - value))

	for i := 0; i < len(intList); i++ {
		if diff > math.Abs(float64(k-intList[i])) {
			diff = math.Abs(float64(k - intList[i]))
			value = intList[i]

		}
	}

	fmt.Println(value)

}
