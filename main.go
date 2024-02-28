package main

import "fmt"

func main() {
	var u *U
	a(u)
}

func a(u interface{}) {
	fmt.Println(u == nil)
}

type U struct {
}

// func MergeSort(list []int) []int {
// 	var length = len(list)
// 	if length < 2 {
// 		return list
// 	}
// 	var mid = length / 2
// 	return merge(MergeSort(list[:mid]), MergeSort(list[mid:]))
// }

// func merge(x, y []int) []int {
// 	var r []int = make([]int, len(x)+len(y))
// 	for i, j := 0, 0; ; {
// 		if i < len(x) && (j == len(y) || x[i] < y[j]) {
// 			r[i+j] = x[i]
// 			i++
// 		} else if j < len(y) {
// 			r[i+j] = y[j]
// 			j++
// 		} else {
// 			break
// 		}
// 	}
// 	return r
// }

// func main() {
// 	var list []int = []int{56, 48, 58, 94, 87, 4, 5, 61, 5, 8, 74, 9, 84, 15, 94, 9, 4, 31, 41, 68, 7, 4, 6, 94, 16, 9, 8, 4}
// 	fmt.Println(MergeSort(list))
// 	// fmt.Println(list)

// 	sort.Ints(list)
// 	fmt.Println(list)
// }
