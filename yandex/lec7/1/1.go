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
	nm := toInt(sc)
	n, m := nm[0], nm[1]
	var events [][]int
	for i := 0; i < m; i++ {
		se := toInt(sc)
		start, end := se[0], se[1]
		events = append(events, []int{start, 1})
		events = append(events, []int{end, 2})
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0] || (events[i][0] == events[j][0] && events[i][1] < events[j][1])
	})

	watchCounter := 0
	part := 0
	unwatched := 0
	for _, event := range events {
		if watchCounter == 0 {
			unwatched += event[0] - part
		}
		if event[1] == 1 {
			watchCounter++
		}
		if event[1] == 2 {
			watchCounter--
		}
		if watchCounter == 0 {
			part = event[0] + 1
		}
	}

	unwatched += n - part
	fmt.Println(unwatched)

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
