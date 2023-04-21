package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	set := make(map[string]bool)

	for _, item := range list {
		if _, ok := set[item]; !ok && item != "" {
			set[item] = true
		}
	}

	fmt.Println(len(set))

}
