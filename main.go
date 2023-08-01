package main

import "fmt"

func main() {
	arr := []int{0, 3, -3, 4, -1}
	fmt.Printf("TwoSumBinarySearch(arr, -1): %v\n", TwoSumBinarySearch(arr, -1))
}

func TwoSumBinarySearch(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{}
	}

	// 使用 `二分查找` 查找 y
	indexByBinarySearch := func(nums []int, num int) int {
		startIndex := 0
		endIndex := len(nums) - 1
		for startIndex <= endIndex {
			i := (startIndex + endIndex) / 2
			if nums[i] == num {
				return i
			}
			if nums[i] > num {
				endIndex = i - 1
			} else {
				startIndex = i + 1
			}
		}
		return -1
	}

	// 不断改变固定的 x
	for i := 0; i < len(nums); i++ {
		// 使用 `线性查找` 查找 y
		j := indexByBinarySearch(nums, target-nums[i])
		if j != i && j != -1 {
			return []int{i, j}
		}
	}
	return []int{}
}
