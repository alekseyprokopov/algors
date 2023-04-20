package main

import "fmt"

func main() {
	var tRoom, tCond int
	var config string
	fmt.Scanf("%d %d\n", &tRoom, &tCond)
	fmt.Scanf("%s\n", &config)

	switch config {
	case "freeze":
		if tRoom <= tCond {
			fmt.Println(tRoom)
		} else {
			fmt.Println(tCond)
		}

	case "heat":
		if tRoom <= tCond {
			fmt.Println(tCond)
		} else {
			fmt.Println(tRoom)
		}
	case "auto":
		fmt.Println(tCond)
	case "fan":
		fmt.Println(tRoom)
	}

}
