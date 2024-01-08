package leetcode

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 *
 * https://leetcode.cn/problems/integer-break/description/
 *
 * algorithms
 * Medium (62.72%)
 * Likes:    1329
 * Dislikes: 0
 * Total Accepted:    300K
 * Total Submissions: 478.1K
 * Testcase Example:  '2'
 *
 * 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
 *
 * 返回 你可以获得的最大乘积 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: n = 2
 * 输出: 1
 * 解释: 2 = 1 + 1, 1 × 1 = 1。
 *
 * 示例 2:
 *
 *
 * 输入: n = 10
 * 输出: 36
 * 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
 *
 *
 *
 * 提示:
 *
 *
 * 2 <= n <= 58
 *
 *
 */

// @lc code=start
func integerBreak(n int) int {
	max := func(a, b, c int) int {
		if a >= b && a >= c {
			return a
		}
		if b >= a && b >= c {
			return b
		}
		if c >= a && c >= b {
			return c
		}
		return 0
	}

	dp := make([]int, n+1)
	dp[2] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*(i-j), j*dp[i-j])
		}
	}
	return dp[n]
}

// @lc code=end

// 使用动态规划求解
// 1. 确定 dp[]: dp[i] 表示分拆数字i，可以得到的最大乘积为dp[i]。
// 2. 确定递推公式: dp[i] = max(dp[i], j*(j-i), j*dp[i-j])。假设拆成两个数，一个数为 j，则另一个数为 i-j，乘积为 j*(i-j)；再求 i-j 的最大乘积，即为 dp[i-j]
// 3. 确定 dp[] 初始化方式: dp[2]=1，dp[3]=2
// 4. 确定遍历顺序: 已知 dp[2]、dp[3] 求 dp[i]，故为顺序遍历
// 5. 举例推导
func integerBreak1(n int) int {
	max := func(a, b, c int) int {
		if a >= b && a >= c {
			return a
		}
		if b >= a && b >= c {
			return b
		}
		if c >= a && c >= b {
			return c
		}
		return 0
	}

	dp := make([]int, n)
	dp[2], dp[3] = 1, 2

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*(i-j), j*dp[i-j])
		}
	}
	return dp[n]
}
