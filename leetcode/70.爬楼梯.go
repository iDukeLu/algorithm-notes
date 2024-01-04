package leetcode

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode.cn/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (54.31%)
 * Likes:    3405
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 2.5M
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 2
 * 输出：2
 * 解释：有两种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶
 * 2. 2 阶
 *
 * 示例 2：
 *
 *
 * 输入：n = 3
 * 输出：3
 * 解释：有三种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶 + 1 阶
 * 2. 1 阶 + 2 阶
 * 3. 2 阶 + 1 阶
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 45
 *
 *
 */

// @lc code=start
func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// @lc code=end

// 1 1 -1
// 2 1+1 2 -2
// 3 1+1+1 2+1 1+2 -3: 一层走2步+二层走1步=1+2=3
// 4 1+1+1+1 2+2 1+2+1 1+1+2 2+1+1 -5: 二层走2步+三层走1步=2+3=5
//
// 使用动态规划求解
// 1. 确定 dp[] 及下标含义: dp[i]表示第i阶有几种爬楼梯的方法
// 2. 确定递推公式: dp[i]=dp[i-1]+dp[i-2]
// 3. 确定 dp[] 如何初始化: dp[1]=1、dp[2]=2
// 4. 确定遍历顺序: 已知 dp[1]、dp[2] 求 dp[i]，故为顺序遍历
// 5. 举例推导 db[]: 当 i=5 时，dp[] 为 1、2、3、5、8、13
func climbStairs1(n int) int {
	if n < 3 {
		return n
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
