package main

import (
	"fmt"

	"github.com/iDukeLu/algorithm-notes/search"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("search.BinarySearch(arr, 9): %v\n", search.BinarySearch(arr, 9))
}
