package main

import "fmt"

func main() {
	var n int
	dict := make(map[string]string)
	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		var key, value string
		fmt.Scanf("%s %s\n", &key, &value)
		dict[key] = value
		dict[value] = key
	}

	var enter string
	fmt.Scanf("%s\n", &enter)
	fmt.Println(dict[enter])

}
