package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	var anneSet, borisSet = make(map[int]bool), make(map[int]bool)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scanf("%d\n", &k)
		anneSet[k] = true
	}

	for i := 0; i < m; i++ {
		var k int
		fmt.Scanf("%d\n", &k)
		borisSet[k] = true
	}

	var in, anneOut, borisOut []int

	for key, _ := range anneSet {
		if _, ok := borisSet[key]; ok {
			in = append(in, key)
		} else {
			anneOut = append(anneOut, key)
		}
	}

	for key, _ := range borisSet {
		if _, ok := anneSet[key]; !ok {
			borisOut = append(borisOut, key)
		}
	}
	printRes(in)
	printRes(anneOut)
	printRes(borisOut)

}

func printRes(data []int) {
	sort.Ints(data)
	fmt.Println(len(data))
	resIn := ""
	for _, item := range data {
		resIn += " " + strconv.Itoa(item)
	}
	fmt.Println(strings.TrimLeft(resIn, " "))
}
