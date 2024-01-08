package leetcode

/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 *
 * https://leetcode.cn/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (67.08%)
 * Likes:    1277
 * Dislikes: 0
 * Total Accepted:    311.6K
 * Total Submissions: 464.3K
 * Testcase Example:  '"abc"'
 *
 * 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
 *
 * 回文字符串 是正着读和倒过来读一样的字符串。
 *
 * 子字符串 是字符串中的由连续字符组成的一个序列。
 *
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由小写英文字母组成
 *
 *
 */

// @lc code=start
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	var result int
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					result++
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					result++
				}
			}
		}
	}

	return result
}

// @lc code=end

// 使用动态规划求解
//  1. 确定 dp[] 含义：dp[i][j] 表示区间范围[i,j] （注意是左闭右闭）的子串是否是回文子串，如果是dp[i][j]为true，否则为false。
//  2. 确定递推公式：
//     整体上是两种，就是s[i]与s[j]相等，s[i]与s[j]不相等这两种。
//     当s[i]与s[j]不相等，那没啥好说的了，dp[i][j]一定是false。
//     当s[i]与s[j]相等时，这就复杂一些了，有如下三种情况
//     情况一：下标i 与 j相同，同一个字符例如a，当然是回文子串
//     情况二：下标i 与 j相差为1，例如aa，也是回文子串
//     情况三：下标：i 与 j相差大于1的时候，例如cabac，此时s[i]与s[j]已经相同了，我们看i到j区间是不是回文子串就看aba是不是回文就可以了，那么aba的区间就是 i+1 与 j-1区间，这个区间是不是回文就看dp[i + 1][j - 1]是否为true。
//  3. 确定 dp[] 如何初始化：dp[i][j]初始化为false
//  4. 确定遍历顺序：由 dp[i+1][j-1] 推导出 dp[i][j]，故 i 逆序遍历，j 顺序遍历
//  5. 举例推导
func countSubstrings1(s string) int {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	var result int
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 { // i、j 相等时，字符为 a；i、j 相差 1 时，字符为 aa
					dp[i][j] = true
					result++
				} else if dp[i+1][j-1] { // i、j 相差大于 1 时，判断内部是否为回文字符串
					dp[i][j] = true
					result++
				}
			}
		}
	}

	return result
}
