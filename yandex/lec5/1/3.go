package main

import (
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	var nShirts int
	fmt.Fscan(file, &nShirts)
	//shirts := make([]int, nShirts)

	var test string
	fmt.Fscanln(file, &test)
	fmt.Println(test)

	//
	//i := 0 //index of shirt
	//j := 0 //index of pants
	//diff := Abs(shirts[i] - pants[j])
	//pair := [2]int{shirts[i], pants[j]}
	//
	//for i < nShirts && j < nPants && diff != 0 {
	//	nowdiff := Abs(shirts[i] - pants[j])
	//	if nowdiff < diff {
	//		diff = nowdiff
	//		pair[0] = shirts[i]
	//		pair[1] = pants[j]
	//	}
	//
	//	if shirts[i] > pants[j] {
	//		j++
	//	} else {
	//		i++
	//	}
	//
	//}
	//fmt.Println(pair[0], pair[1])

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
