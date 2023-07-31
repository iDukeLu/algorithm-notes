/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode.cn/problems/two-sum/description/
 *
 * algorithms
 * Easy (52.93%)
 * Likes:    17391
 * Dislikes: 0
 * Total Accepted:    4.6M
 * Total Submissions: 8.8M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
 *
 * 你可以按任意顺序返回答案。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,7,11,15], target = 9
 * 输出：[0,1]
 * 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,2,4], target = 6
 * 输出：[1,2]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [3,3], target = 6
 * 输出：[0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * 只会存在一个有效答案
 *
 *
 *
 *
 * 进阶：你可以想出一个时间复杂度小于 O(n^2) 的算法吗？
 *
 */

package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{}
	}

	// 构造一个 Map，key: num，value: index
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	// 不断改变固定的 x
	for i := 0; i < len(nums); i++ {
		// 使用 `Map` 查找 y
		if j, ok := m[target-nums[i]]; ok {
			if j != i {
				return []int{i, j}
			}
		}

	}
	return []int{}
}

// @lc code=end

// 问题理解：
// 		x + y = target，返回满足条件的 x，y 的坐标
// 问题分析：
// 		第一眼看存在 x、y 两个未知数需要解，实际我们可以先固定 x 再查找符合条件的 y，
// 		问题就回到了不断改变固定的 x，然后查找 y，这样就回到了 '数组的查找问题'
// 求解思考：
//		数组的查找问题：目前，我的知识库中只有 `线性查找` 和 `二分查找` 两种方式
// 		不断改变固定的 x，复杂度为 n；
//			- 使用 `线性查找` 查找 y，复杂度为 n，整体复杂度为 n^2
//			- 使用 `二分查找` 查找 y，复杂度为logn，整体复杂度为 nlogn

// 求解1：使用线性查找求解，也就是所谓的暴力查找，整体时间复杂度为 n^2
func TwoSumLinearSearch(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{}
	}

	// 使用 `线性查找` 查找 y
	indexByLinearSearch := func(nums []int, num int) int {
		for i := 0; i < len(nums); i++ {
			if nums[i] == num {
				return i
			}
		}
		return -1
	}

	// 不断改变固定的 x
	for i := 0; i < len(nums); i++ {
		// 使用 `线性查找` 查找 y
		j := indexByLinearSearch(nums, target-nums[i])
		if j != i && j != -1 {
			return []int{i, j}
		}
	}
	return []int{}
}

// 求解2：使用二分查找求解，整体时间复杂度为 nlogn
func TwoSumBinarySearch(nums []int, target int) []int {

	return []int{}
}

// 求解思考：
// 		剩下两个思考方向：
// 			1. 如何降低 `不断改变固定的 x` 的复杂度

// 求解3：使用 Map 查找求解，整体时间复杂度为 n
func TwoSumMapSearch(nums []int, target int) []int {
	if len(nums) <= 0 {
		return []int{}
	}

	// 构造一个 Map，key: num，value: index
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	// 不断改变固定的 x
	for i := 0; i < len(nums); i++ {
		// 使用 `Map` 查找 y
		if j, ok := m[target-nums[i]]; ok {
			if j != i {
				return []int{i, j}
			}
		}

	}
	return []int{}
}
