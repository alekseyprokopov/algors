package main

//
//import (
//	"fmt"
//)
//
//func main() {
//	var nShirts, nPants int
//
//	fmt.Scanf("%d", &nShirts)
//	shirts := make([]int, nShirts)
//
//	for i := 0; i < nShirts; i++ {
//		fmt.Scanf("%d", &shirts[i])
//	}
//
//	fmt.Scanf("%d", &nPants)
//	pants := make([]int, nPants)
//
//	for i := 0; i < nPants; i++ {
//		fmt.Scanf("%d", &pants[i])
//	}
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
