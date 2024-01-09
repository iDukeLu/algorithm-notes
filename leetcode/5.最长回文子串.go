package leetcode

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (37.94%)
 * Likes:    7006
 * Dislikes: 0
 * Total Accepted:    1.6M
 * Total Submissions: 4.2M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 * 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	var result string
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 || dp[i+1][j-1] {
					dp[i][j] = true
					if j+1-i >= len(result) {
						result = s[i : j+1]
					}
				}
			}
		}
	}

	return result
}

// @lc code=end

// 使用动态规划求解：思路类似 647 题
func longestPalindrome1(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	var result string
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					if j+1-i >= len(result) {
						result = s[i : j+1]
					}
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					if j+1-i >= len(result) {
						result = s[i : j+1]
					}
				}
			}
		}
	}

	return result
}

// GPT 优化后版本
// 1. 初始化优化：目前的代码中，对 dp 数组的初始化是通过两层循环完成的。这是不必要的，因为你只需要初始化对角线元素（单个字符总是回文），以及一些你实际会使用到的元素。初始化可以更简洁。
// 2. 判断条件简化：在内层循环中，if 判断条件可以合并，以减少代码的嵌套层次。
// 3. 结果更新优化：目前的代码在每次找到一个更长的回文子串时都会更新 result。你可以通过在循环开始前记录最长回文子串的起始位置和长度，而不是直接记录子串本身，来减少字符串操作，提高效率。
// 4. 代码可读性：添加一些注释，说明每一部分代码的作用，这对于理解和维护代码都是有帮助的。
func longestPalindrome_gpt(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true // 单个字符是回文
	}

	start, maxLen := 0, 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				// 仅有两个字符或子串是回文
				if j-i == 1 || dp[i+1][j-1] {
					dp[i][j] = true
					if j-i+1 > maxLen {
						start = i
						maxLen = j - i + 1
					}
				}
			}
		}
	}

	return s[start : start+maxLen]
}
