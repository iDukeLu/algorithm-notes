package sort

// 维基百科：
// 归并排序（merge sort 或 mergesort）：是基于 *归并操作* 的一种有效的排序算法，
// 该算法是采用 *分治法* 的一个非常典型的应用，且各层分治递归可以同时进行，1945 年由约 *翰·冯·诺伊曼* 首次提出
// 归并操作（merge）：也叫归并算法，指的是_将两个已经排序的序列合并成一个排序序列_的操作。
//
// 归并排序：不断使用归并操作进行排序的算法。
// 归并操作：将两个排序序列合并成一个排序序列，如何合并：不断比较两个序列的首元素，并将较小的元素依次放入合并的新集合中
//
// 复杂度：
// - 平均时间复杂度：O(nlogn)
// - 最坏时间复杂度：O(nlogn)
// - 最优时间复杂度：O(nlogn)
// - 空间复杂度：O(n)
//
// 实现方式：
// - 递归法（Top-down）
// - 迭代法（Bottom-up）
func MergeSort[T Number](arr []T) []T {
	merge := func(arr1 []T, arr2 []T) []T {
		merged := make([]T, 0, len(arr1)+len(arr2))
		i, j := 0, 0

		for i < len(arr1) && j < len(arr2) {
			if arr1[i] < arr2[j] {
				merged = append(merged, arr1[i])
				i++
			} else {
				merged = append(merged, arr2[j])
				j++
			}
		}

		// 如果 arr1 还有剩余元素，直接追加
		merged = append(merged, arr1[i:]...)
		// 如果 arr2 还有剩余元素，直接追加
		merged = append(merged, arr2[j:]...)

		return merged
	}

	var length = len(arr)
	if length < 2 {
		return arr
	}
	return merge(MergeSort(arr[:length/2]), MergeSort(arr[length/2:]))
}

// 根据归并排序的思路，推到一遍
func MergeSortStep[T Number](arr []T) []T {

	// 定义归并方法
	merge := func(arr1 []T, arr2 []T) []T {
		var arr []T
		for i, j := 0, 0; ; {
			if i < len(arr1) { // 数组1: 存在元素
				val1 := arr1[i]
				if j < len(arr2) { // 数组2: 存在元素
					val2 := arr2[j]
					if val1 < val2 {
						arr = append(arr, val1)
						i++
					} else {
						arr = append(arr, val2)
						j++
					}
				} else { // 数组2: 不存在元素
					arr = append(arr, val1)
					i++
				}
			} else { // 数组1: 不存在元素
				if j < len(arr2) { // 数组2: 存在元素
					val2 := arr2[j]
					arr = append(arr, val2)
					j++
				} else { // 数组2: 不存在元素
					// 此时，数组1 和 数组2 均不存在元素直接结束
					break
				}
			}
		}
		return arr
	}

	length := len(arr)
	if length < 2 { // 数组长度 < 2，直接返回
		return arr
	}
	// 数组长度 > 2，继续切分合并
	return merge(MergeSortStep(arr[:length/2]), MergeSortStep(arr[length/2:]))
}
