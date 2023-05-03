package main

import "fmt"

// найти наибольний убывающий отрезок
func main() {
	//arr := []int{6, 1, 5, 2, 5, 7, 6, 5, 4, 8, 1, 2}
	//arr := []int{6, 5, 4, 3, 2, 1}
	arr := []int{6, 6}
	fmt.Println(biggest(arr))
}

func biggest(arr []int) (int, int) {
	maxCounter := 0
	counter := 0
	right_index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] <= arr[i] {
			counter++
		} else {
			counter = 0
		}
		if counter > maxCounter {
			maxCounter = counter
			right_index = i
		}
	}

	return right_index - maxCounter, right_index
}
