package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	START = -1
	POINT = 0
	END   = 1
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	nm := toInt(sc)
	n, _ := nm[0], nm[1]
	var points []int
	var events [][]int

	for i := 0; i < n; i++ {
		event := toInt(sc)
		start, end := event[0], event[1]
		events = append(events, []int{start, START}, []int{end, END})
		//events = append(events, []int{end, END})
	}
	points = toInt(sc)
	for _, point := range points {
		events = append(events, []int{point, POINT})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0] ||
			(events[i][0] == events[j][0] && events[i][1] < events[j][1])
	})

	counterNow := 0
	var dict = make(map[int]int)

	for _, event := range events {
		if event[1] == START {
			counterNow++
		}
		if event[1] == POINT {
			dict[event[0]] = counterNow
		}
		if event[1] == END {
			counterNow--
		}
	}
	var res []string
	for _, point := range points {
		res = append(res, strconv.Itoa(dict[point]))
	}
	fmt.Println(strings.Join(res, " "))
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
