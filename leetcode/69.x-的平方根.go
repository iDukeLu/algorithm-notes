package leetcode

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode.cn/problems/sqrtx/description/
 *
 * algorithms
 * Easy (38.46%)
 * Likes:    1497
 * Dislikes: 0
 * Total Accepted:    838.8K
 * Total Submissions: 2.2M
 * Testcase Example:  '4'
 *
 * 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
 *
 * 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
 *
 * 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：x = 4
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：x = 8
 * 输出：2
 * 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= x <= 2^31 - 1
 *
 *
 */

// @lc code=start
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 2, x/2
	var mid int

	for left <= right {
		mid = left + (right-left)/2
		squared := mid * mid

		if squared == x {
			return mid
		}

		if squared < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

// @lc code=end

// 使用二分查找法。这种方法适用于查找平方根的整数部分
// 基本思想是在 0 到 x 之间查找一个数 mid，使得 mid * mid 最接近但不大于 x。
