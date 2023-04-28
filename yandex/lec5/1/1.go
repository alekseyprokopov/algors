package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//
//	file, _ := os.Open("input.txt")
//	defer file.Close()
//	scan := bufio.NewScanner(file)
//	scan.Split(bufio.ScanLines)
//	var nShirts, nPants int
//
//	scan.Scan()
//	nShirts, _ = strconv.Atoi(scan.Text())
//	shirts := make([]int, nShirts)
//
//	scan.Scan()
//	for i, item := range strings.Fields(scan.Text()) {
//		shirts[i], _ = strconv.Atoi(item)
//	}
//
//	scan.Scan()
//	nPants, _ = strconv.Atoi(scan.Text())
//	pants := make([]int, nPants)
//
//	scan.Scan()
//	for i, item := range strings.Fields(scan.Text()) {
//		pants[i], _ = strconv.Atoi(item)
//	}
//
//	i := 0 //index of shirt
//	j := 0 //index of pants
//	diff := Abs(shirts[i] - pants[j])
//	pair := [2]int{shirts[i], pants[j]}
//
//	for i < nShirts && j < nPants && diff != 0 {
//		nowdiff := Abs(shirts[i] - pants[j])
//		if nowdiff < diff {
//			diff = nowdiff
//			pair[0] = shirts[i]
//			pair[1] = pants[j]
//		}
//
//		if shirts[i] > pants[j] {
//			j++
//		} else {
//			i++
//		}
//
//	}
//	fmt.Println(pair[0], pair[1])
//
//}
//
//func Abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
