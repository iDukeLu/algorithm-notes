package leetcode

/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 *
 * https://leetcode.cn/problems/binary-search/description/
 *
 * algorithms
 * Easy (54.90%)
 * Likes:    1514
 * Dislikes: 0
 * Total Accepted:    1.1M
 * Total Submissions: 2.1M
 * Testcase Example:  '[-1,0,3,5,9,12]\n9'
 *
 * 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的
 * target，如果目标值存在返回下标，否则返回 -1。
 *
 *
 * 示例 1:
 *
 * 输入: nums = [-1,0,3,5,9,12], target = 9
 * 输出: 4
 * 解释: 9 出现在 nums 中并且下标为 4
 *
 *
 * 示例 2:
 *
 * 输入: nums = [-1,0,3,5,9,12], target = 2
 * 输出: -1
 * 解释: 2 不存在 nums 中因此返回 -1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 你可以假设 nums 中的所有元素是不重复的。
 * n 将在 [1, 10000]之间。
 * nums 的每个元素都将在 [-9999, 9999]之间。
 *
 *
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// @lc code=end

// 第一版：思路对了，但是每次传入的是二分后的子数组，导致返回的索引也是子数组的索引
func search1(nums []int, target int) int {
	l := len(nums)
	i := l / 2

	if nums[i] == target {
		return i
	}

	if l == 1 && nums[i] != target {
		return -1
	}
	if nums[i] > target {
		return search(nums[:i], target)
	} else {
		return search(nums[i:], target)
	}
}

// 根据第一版修改的实现，但是边界条件判断的不是很优雅
func search2(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for i := (s + e + 1) / 2; ; i = (s + e + 1) / 2 {
		v := nums[i]
		if s == e && v != target {
			return -1
		}
		if v == target {
			return i
		}
		if v > target {
			e = i - 1
			if e < 0 {
				e = 0
			}
		} else {
			s = i + 1
			if s > len(nums)-1 {
				s = len(nums) - 1
			}
		}
	}
}

// GPT 优化版本
// 1. 避免无限循环：原代码中的 for 循环缺少一个明确的退出条件，这可能会导致无限循环。通常，在二分查找中，当 left 大于 end 时，表示没有找到目标元素，应该退出循环。
// 2. 简化边界更新逻辑：在更新 left 和 end 指针时，无需额外的边界检查。这些检查是多余的，因为在二分查找的逻辑中，left 和 end 总是会保持在有效范围内。
// 3. 循环条件和退出逻辑：使用 while 或 for 循环，当 left 小于等于 end 时持续查找，如果找到目标，则返回索引；否则，在循环结束后返回 -1。
func search_gpt(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
