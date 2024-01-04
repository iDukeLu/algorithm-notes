package leetcode

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 *
 * https://leetcode.cn/problems/min-cost-climbing-stairs/description/
 *
 * algorithms
 * Easy (65.97%)
 * Likes:    1427
 * Dislikes: 0
 * Total Accepted:    389.2K
 * Total Submissions: 589.8K
 * Testcase Example:  '[10,15,20]'
 *
 * 给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
 *
 * 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
 *
 * 请你计算并返回达到楼梯顶部的最低花费。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：cost = [10,15,20]
 * 输出：15
 * 解释：你将从下标为 1 的台阶开始。
 * - 支付 15 ，向上爬两个台阶，到达楼梯顶部。
 * 总花费为 15 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：cost = [1,100,1,1,1,100,1,1,100,1]
 * 输出：6
 * 解释：你将从下标为 0 的台阶开始。
 * - 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
 * - 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
 * - 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
 * - 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
 * - 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
 * - 支付 1 ，向上爬一个台阶，到达楼梯顶部。
 * 总花费为 6 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= cost.length <= 1000
 * 0 <= cost[i] <= 999
 *
 *
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {

	min := func(i1, i2 int) int {
		if i1 < i2 {
			return i1
		}
		return i2
	}

	dp := make([]int, len(cost)+1)
	dp[0], dp[1] = 0, 0

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

// @lc code=end

// [1,100,1,1,1,100,1,1,100,1]
//  0  1  2 3 4  5  6 7  8  9
//  0  0  1 2 2  3  3 4  4  5 6
//
// 1. 确定 dp[] 和下标含义: 第i阶需要花费的总费用
// 2. 确定递推公式: dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
// 3. 确定 dp[] 如何初始化: dp[0]=0、dp[1]=0，第 0、1 阶均花费为 0，因为 "一旦你支付此费用，即可选择向上爬一个或者两个台阶。"，但在第 0、1 阶时还未支付
// 4. 确定遍历顺序: 已知 dp[0]、dp[1] 求 dp[i]，故为顺序遍历
// 5. 举例推导 db[]:
