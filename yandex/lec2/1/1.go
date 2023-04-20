package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	list := strings.Split(scanner.Text(), " ")
	flag := true

	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			flag = false
			break
		}
	}

	if flag {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}
