package leetcode

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode.cn/problems/unique-paths/description/
 *
 * algorithms
 * Medium (67.92%)
 * Likes:    1963
 * Dislikes: 0
 * Total Accepted:    713K
 * Total Submissions: 1M
 * Testcase Example:  '3\n7'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
 *
 * 问总共有多少条不同的路径？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：m = 3, n = 7
 * 输出：28
 *
 * 示例 2：
 *
 *
 * 输入：m = 3, n = 2
 * 输出：3
 * 解释：
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向下
 *
 *
 * 示例 3：
 *
 *
 * 输入：m = 7, n = 3
 * 输出：28
 *
 *
 * 示例 4：
 *
 *
 * 输入：m = 3, n = 3
 * 输出：6
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m, n <= 100
 * 题目数据保证答案小于等于 2 * 10^9
 *
 *
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][1] = 1
	}
	for i := 1; i <= n; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}

// @lc code=end

// m/n| 1  2  3
// ————————————
//
//	1 | 1  1  1
//	2 | 1  2  3
//	3 | 1  3  6
//
// 使用动态规划求解
// 1. 确定 dp[] 含义: d[m][n] 表示 m x n 网格时，有多少条不同的路径
// 2. 确定递推公式: d[m][n] = d[m-1][n] + d[m][n-1]
// 3. 确定 dp[] 如何初始化: dp[m][1]=1, dp[1][n]=1
// 4. 确定遍历顺序: 已知 dp[m][1]=1, dp[1][n]=1 求 d[m][n]，故为顺序遍历
// 5. 举例推导:
func uniquePaths1(m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][1] = 1
	}

	for i := 1; i <= n; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}
