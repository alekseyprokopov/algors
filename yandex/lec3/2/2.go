package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list1 := strings.Split(sc.Text(), " ")
	sc.Scan()
	list2 := strings.Split(sc.Text(), " ")

	dict1 := make(map[string]bool)

	var res []string
	for _, item := range list1 {
		dict1[item] = true
	}

	for _, item := range list2 {
		if _, ok := dict1[item]; ok {
			res = append(res, item)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		v1, _ := strconv.Atoi(res[i])
		v2, _ := strconv.Atoi(res[j])
		return v1 < v2
	})
	fmt.Println(strings.Join(res, " "))
}
