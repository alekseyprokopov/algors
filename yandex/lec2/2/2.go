package main

import (
	"fmt"
)

func main() {
	list := []int{}

	for {
		var enter int
		fmt.Scanf("%d\n", &enter)
		if enter == -2000000000 {
			break
		}
		list = append(list, enter)
	}
	constant, asc, wAsc, desc, wDesc, rand := true, true, true, true, true, true

	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			asc = false

			if list[i] == list[i-1] {
				desc = false
			} else {
				constant = false
				wAsc = false
			}

		} else {
			desc = false

			if list[i] == list[i-1] {
				asc = false
			} else {
				constant = false
				wDesc = false
			}

		}
	}

	if constant {
		fmt.Println("CONSTANT")
	} else if asc {
		fmt.Println("ASCENDING")
	} else if wAsc {
		fmt.Println("WEAKLY ASCENDING")
	} else if desc {
		fmt.Println("DESCENDING")
	} else if wDesc {
		fmt.Println("WEAKLY DESCENDING")
	} else if rand {
		fmt.Println("RANDOM")
	}

}
