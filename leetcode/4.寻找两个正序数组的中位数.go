package leetcode

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (41.72%)
 * Likes:    6943
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 2.5M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
 *
 * 算法的时间复杂度应该为 O(log (m+n)) 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 *
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	merged := make([]int, 0, length)
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	merged = append(merged, nums1[i:]...)
	merged = append(merged, nums2[j:]...)

	if length%2 == 0 {
		return float64(merged[length/2-1]+merged[length/2]) / 2
	}
	return float64(merged[length/2])
}

// @lc code=end

// 关键字：
// - 正序：说明已经排好序
// - 中位数：一个数值集合中的一个中间值

// 暴力求解：先合并两个数组，然后返回中位数
// 但是时间复杂度为：O(m+n) 不满足题目要求
// 进一步降低时间复杂度需使用二分法
func force_findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	merged := make([]int, 0, length)
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	merged = append(merged, nums1[i:]...)
	merged = append(merged, nums2[j:]...)

	if length%2 == 0 {
		return float64(merged[length/2-1]+merged[length/2]) / 2
	}
	return float64(merged[length/2])
}
