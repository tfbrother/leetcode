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
// 在滑动窗口中做记录，采用查找表来做记录
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

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Longest Substring Without Repeating Characters.
*/
// 用数组替代map，但是不能处理包含中文的情况
// 查找表+滑动窗口技术
func lengthOfLongestSubstring2(s string) int {
	l, r, length, max := 0, 0, len(s), 0
	if length <= 1 {
		return length
	}

	var charIndex [128]int
	// 初始化为-1很重要，因为默认值是0，索引也有可能为0。为了下面和l比较的时候区别。
	for i := 0; i < 128; i++ {
		charIndex[i] = -1
	}

	for r < length {
		// s[r]之前出现过，并且在窗口的内
		if charIndex[s[r]] >= l {
			l = charIndex[s[r]] + 1
		}

		charIndex[s[r]] = r
		if max < r-l+1 {
			max = r - l + 1
		}
		r++
	}

	return max
}
