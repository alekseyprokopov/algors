package main

import "fmt"

//массив числел найти пару, максимально близкую к К

func main() {
	arr := []int{1, 2, 3, 6, 6} //3,4
	k := 12
	fmt.Println(arr, k)
	fmt.Println(closest(arr, k))
}

func closest(arr []int, target int) (int, int) {
	l := 0
	r := len(arr) - 1
	best_l, best_r := l, r

	for l < r {
		best_sum := Abs(arr[best_l] + arr[best_r] - target)
		now_sum := Abs(arr[l] + arr[r] - target)
		if now_sum < best_sum {
			best_l, best_r = l, r
		}

		if arr[best_l]+arr[best_r] >= target {
			fmt.Println("right")
			r--
		} else {
			fmt.Println("left")
			l++
		}
	}
	return best_l, best_r
}

func Abs(k int) int {
	if k >= 0 {
		return k
	}
	return -k
}
