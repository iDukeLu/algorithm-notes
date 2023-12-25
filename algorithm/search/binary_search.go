package search

// 二分查找法
// TODO，补充定义和思想
func BinarySearch(nums []int, num int) int {
	return BinarySearch4(nums, num)
}

// 最终代码：BinarySearch3、BinarySearch4

// 第 1 次推导
func BinarySearchStep1(nums []int, num int) int {
	// 第 1 轮查找
	i := len(nums) / 2
	if nums[i] == num {
		return i
	}
	if nums[i] > num {
		nums = nums[:i]
	}
	if nums[i] < num {
		nums = nums[i:]
	}

	// 第 2 轮查找
	i = len(nums) / 2
	if nums[i] == num {
		return i
	}
	if nums[i] > num {
		nums = nums[:i]
	}
	if nums[i] < num {
		nums = nums[i:]
	}
	// 第 n 轮查找...

	// 以上方式在尝试运行后，只能查找是否存在对应的元素，不能返回查找元素的索引
	// 因为 nums 数组被不断重新赋值，导致原有数组的索引丢失
	return 0
}

// 第 1 次结论：无法正常返回查找元素的索引
// 因为 nums 数组被不断重新赋值，导致原有数组的索引丢失
func BinarySearch1(nums []int, num int) int {
	i := len(nums) / 2
	if nums[i] == num {
		return i
	}
	if nums[i] > num {
		nums = nums[:i]
	}
	if nums[i] < num {
		nums = nums[i:]
	}
	return BinarySearch1(nums, num)
}

// 第 2 次推导
func BinarySearchStep2(nums []int, num int) int {
	// 接下来换种方式思考，不对原数组变量 nums 进行赋值，而是控制循环的范围

	startIndex := 0
	endIndex := len(nums) - 1

	// 简单推导一下，可以得到 中间索引 = (开始索引 + 结素索引)/2

	// 第 1 轮查找
	i := (startIndex + endIndex) / 2
	if nums[i] == num {
		return i
	}
	if nums[i] > num {
		endIndex = i
	}
	if nums[i] < num {
		startIndex = i
	}

	// 第 2 轮查找
	i = (startIndex + endIndex) / 2
	if nums[i] == num {
		return i
	}
	if nums[i] > num {
		endIndex = i
	}
	if nums[i] < num {
		startIndex = i
	}
	// 第 n 轮查找...

	return 0
}

// 第 2 次结论：
//  1. startIndex = i 和 endIndex = i 无法很好的缩小范围，可能进入死循环，
//     比如 [1,2] 在查找 1 时，会导致 ，endIndex 无限等于 1
//  2. 要单独判断 startIndex == endIndex 的情况，看起来很奇怪
func BinarySearch2(nums []int, num int) int {
	startIndex := 0
	endIndex := len(nums) - 1
	for {
		i := (startIndex + endIndex) / 2
		if startIndex == endIndex {
			if nums[i] == num {
				return startIndex
			} else {
				return -1
			}
		}

		if nums[i] == num {
			return i
		}
		if nums[i] > num {
			endIndex = i
		}
		if nums[i] < num {
			startIndex = i
		}
	}
}

// 第 3 次结论：
// 改进了判断条件，不用单独处理 startIndex == endIndex 的特殊情况
// endIndex = i - 1 和 startIndex = i + 1 可以有效缩小范围
func BinarySearch3(nums []int, num int) int {
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

// 第 4 次结论：使用递归替换 for 循环
func BinarySearch4(nums []int, num int) int {
	return binarySearch4(nums, num, 0, len(nums)-1)
}

func binarySearch4(nums []int, num int, startIndex, endIndex int) int {
	if startIndex <= endIndex {
		i := (startIndex + endIndex) / 2
		if nums[i] == num {
			return i
		}
		if nums[i] > num {
			endIndex = i - 1
		} else {
			startIndex = i + 1
		}
		return binarySearch4(nums, num, startIndex, endIndex)
	}
	return -1
}
