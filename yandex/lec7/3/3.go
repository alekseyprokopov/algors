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
	CAN  = 0
	CANT = 1
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	nd := toInt(sc)
	n, d := nd[0], nd[1]
	coords := toInt(sc)
	var events [][]int
	for i := 0; i < n; i++ {
		events = append(events, []int{coords[i], CAN})
		events = append(events, []int{coords[i] + d, CANT})
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0] ||
			(events[i][0] == events[j][0] && events[i][1] < events[j][1])
	})
	max, count := 0, 0

	for _, event := range events {
		fmt.Println("count:", count)
		eventType := event[1]
		if eventType == CAN {
			count++
		} else if eventType == CANT {
			count--
		}

		if count > max {
			max = count
		}
	}
	dict := map[int]int{}
	variant := 1
	for _, event := range events {
		x, evenType := event[0], event[1]
		if evenType == CAN {
			dict[x] = variant
		}
		variant++
		if variant > max {
			variant = 1
		}
	}

	var res []string
	for _, point := range coords {
		res = append(res, strconv.Itoa(dict[point]))
	}

	fmt.Println(max)
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
