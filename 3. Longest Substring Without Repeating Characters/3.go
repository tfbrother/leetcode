package ___Longest_Substring_Without_Repeating_Characters

// see https://leetcode-cn.com/problems//longest-substring-without-repeating-characters/
// 相似题
// 	159. Longest Substring with At Most Two Distinct Charactersg  				Hard
//	340. Longest Substring with At Most K Distinct Characters		Hard
//	992. Subarrays with K Different Integers 	Hard

/*
Runtime: 4 ms, faster than 91.21% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 3.1 MB, less than 56.29% of Go online submissions for Longest Substring Without Repeating Characters.
*/
// 解法一：滑动窗口技术
// 在滑动窗口中做记录
func lengthOfLongestSubstring(s string) int {
	// 初始值的设定很重要，一定要搞清楚每一个的定义
	// i，j表示窗口的左右边界，charIndex存储的是已遍历的字符在s中最后出现的位置，主要用来和i比较的。
	i, j, charIndex, max, find, length, index := 0, 0, make(map[byte]int), 0, false, 0, -1

	for j < len(s) {
		// 前面出现过该字符
		if index, find = charIndex[s[j]]; find {
			if i <= index {
				i = index + 1
			}
		}

		length = j - i + 1
		if max < length {
			max = length
		}

		charIndex[s[j]] = j
		j++
	}

	return max
}
