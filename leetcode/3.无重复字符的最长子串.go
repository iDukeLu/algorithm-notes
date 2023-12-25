package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (39.13%)
 * Likes:    9764
 * Dislikes: 0
 * Total Accepted:    2.6M
 * Total Submissions: 6.5M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var l, r, max int

	for l < len(s) {

		for r < len(s) && m[s[r]] == 0 {
			m[s[r]]++ // 窗口添加元素
			r++       // 右指针移动
		}

		if len(m) > max {
			max = len(m)
		}

		delete(m, s[l]) // 窗口删除元素
		l++             // 左指针移动
	}

	return max
}

// @lc code=end

// 理解：
//		从一个字符串中找出无重复字符的最长子串，返回结果为子串的长度
// 思路：
//		实际也是数组查找的问题
// 		1. 暴力求解（线性查找）：双重循环，找出所有无重复字串（无重复 - 使用 Map 保存字符），循环所有字串找到最长的，复杂度 O(n^2) - 太过暴力，考虑的场景太多，完全是穷举
// 		2. 使用 Map[sting][]int 记录每个字符串的坐标：循环 Map 找出坐标最大差，即最大长度，复杂度  O(n) + O(n^2) - 错误解法：这种只能比较第一个字符和最后一个字符，中间字符没有比较
// 正解：
// 		1. 使用滑动窗口

// 暴力求解（线性查找）：双重循环，找出所有无重复字串，循环所有字串找到最长的，复杂度 O(n^2) + O(n)、
// 非常不推荐：太过暴力，考虑的场景太多，完全是穷举
func force(s string) int {
	if len(s) == 1 {
		return len(s)
	}

	var strs []string

	for i1, v1 := range s {
		str := []rune{v1}
		m := make(map[rune]struct{})
		m[v1] = struct{}{}
		for i2, v2 := range s {
			if i2 > i1 {
				_, ok := m[v2]
				if !ok {
					m[v2] = struct{}{}
					str = append(str, v2)
					if i2 == len(s)-1 {
						strs = append(strs, string(str))
						break
					}
				} else {
					strs = append(strs, string(str))
					break
				}
			}
		}
	}

	var max int
	for _, str := range strs {
		if len(str) > max {
			max = len(str)
		}
	}
	return max
}

// 滑动窗口，求解：使用两个指针表示窗口的边界，通过不断移动改变窗口边界查找无重复的字符串并记录其长度
// 时间复杂度：O(N)、空间复杂度：O(∣Σ∣)
func slidingWindow(s string) int {
	m := make(map[byte]int)
	j, max := 0, 0

	for i := 0; i < len(s); i++ {

		for j < len(s) && m[s[j]] == 0 {
			m[s[j]]++
			j++
		}

		if len(m) > max {
			max = len(m)
		}
		delete(m, s[i])
	}

	return max
}
