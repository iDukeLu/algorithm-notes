package leetcode

/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 *
 * https://leetcode.cn/problems/longest-palindromic-subsequence/description/
 *
 * algorithms
 * Medium (67.11%)
 * Likes:    1154
 * Dislikes: 0
 * Total Accepted:    213.1K
 * Total Submissions: 317.4K
 * Testcase Example:  '"bbbab"'
 *
 * 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
 *
 * 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "bbbab"
 * 输出：4
 * 解释：一个可能的最长回文子序列为 "bbbb" 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出：2
 * 解释：一个可能的最长回文子序列为 "bb" 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由小写英文字母组成
 *
 *
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	size := len(s)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
		dp[i][i] = 1
	}

	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][size-1]
}

// @lc code=end

// cbbd
// i/j   0    1     2       3
// 0   1(c) 1(cb) 2(cbb) 2(cbbd)
// 1    -   1(b)  2(bb)  2(bdb)
// 2    -    -    1(b)   1(bd)
// 3    -    -     -     1(d)
// 使用动态规划求解
//  1. 确定 dp[]：dp[i][j] 表示字符串s在[i, j]范围内最长的回文子序列的长度
//  2. 确定递推公式：
//     s[i] == s[j] 时，dp[i][j] = dp[i+1][j-1] + 2
//     s[i] != s[j] 时，dp[i][j] = max(dp[i+1][j], dp[i][j-1])
//  3. 确定 dp[] 如何初始化：i==j 时，dp[i][j]=1
//  4. 确定遍历顺序：已知 dp[i+1][j-1] 求 dp[i][j]，故 i 顺序遍历，j 逆序遍历
//  5. 举例推导
func longestPalindromeSubseq_dp(s string) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

// GTP 优化版本
// 1. 空间优化：当前实现使用了一个二维数组 dp，实际上可以使用一个一维数组来节省空间。因为在计算 dp[i][j] 时，只需要 dp[i+1][j-1]、dp[i+1][j] 和 dp[i][j-1] 的值。这意味着我们可以用一维数组从右到左更新，每次只保留所需的上一个状态。
// 2. 代码简化：可以将 max 函数的定义内联到代码中，以简化代码结构。这在Go语言中是常见的做法，因为Go的标准库没有提供内置的 max 函数。
// 3. 代码可读性：添加一些注释可以帮助解释代码的意图，这对于维护和理解代码是很有帮助的。
func longestPalindromeSubseq_gpt(s string) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(s)
	if n <= 1 {
		return n
	}

	dp := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		prev := 0
		for j := i + 1; j < n; j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = prev + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = temp
		}
	}

	return dp[n-1]
}
