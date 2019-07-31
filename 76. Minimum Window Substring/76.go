package _6__Minimum_Window_Substring

// https://leetcode-cn.com/problems/minimum-window-substring/
/*
相似题目：
	串联所有单词的子串 	困难
	长度最小的子数组 		中等
	滑动窗口最大值 		困难
	字符串的排列 			中等
	最小区间 			困难
	最小窗口子序列 		困难
*/

/*
Runtime: 24 ms, faster than 52.42% of Go online submissions for Minimum Window Substring.
Memory Usage: 2.9 MB, less than 95.65% of Go online submissions for Minimum Window Substring.
*/
// 滑动窗口算法，和209，438相似。
// 题意理解：
// 	1. 需要考虑大小写否？
// 	2. t中可能包含有重复的字符
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// l,r为滑动窗口的左右边界，width表示窗口的最大值。
	l, r, start, minLen, width, charCount1, charCount2 := 0, 0, 0, len(s)+1, 0, make(map[byte]int), make(map[byte]int)
	for k := 0; k < len(t); k++ {
		charCount2[t[k]]++
	}

	for r < len(s) {
		// fmt.Println("l===", l, ";r===", r)
		if _, find := charCount2[s[r]]; find { // 窗口从j+1开始
			charCount1[s[r]]++

			if charCount1[s[r]] == charCount2[s[r]] {
				width++
			}
		}
		r++

		for width == len(charCount2) {
			if minLen > r-l {
				minLen = r - l
				start = l
			}

			if charCount2[s[l]] > 0 {
				charCount1[s[l]]--
				if charCount1[s[l]] < charCount2[s[l]] {
					width--
				}
			}
			l++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return string(s[start : start+minLen])
}
