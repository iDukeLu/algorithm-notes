package search

// 二分
func BinarySearch(nums []int, num int) int {
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
	return BinarySearch(nums, num)
}

func BinarySearchStep(nums []int, num int) int {

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

	// 以上方式不可行，在尝试运行后，发现 nums 数组会被不断

	return 0
}
