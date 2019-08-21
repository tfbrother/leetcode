package _90

import "strings"

// https://leetcode-cn.com/problems/word-pattern/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Pattern.
Memory Usage: 2.1 MB, less than 25.00% of Go online submissions for Word Pattern.
*/
// 完全匹配，这种双向匹配关系由两个映射map来存储。
// 第一版的实现只存储了单向的匹配，没有考虑到strMap。
func wordPattern(pattern string, str string) bool {
	strs, paMap, strMap := splitStr(str), make(map[byte]string), make(map[string]byte)
	if len(strs) != len(pattern) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		value, ok1 := paMap[pattern[i]]
		key, ok2 := strMap[strs[i]]
		if !ok1 && !ok2 {
			paMap[pattern[i]], strMap[strs[i]] = strs[i], pattern[i]
		} else if value != strs[i] || key != pattern[i] {
			return false
		}
	}

	return true
}

// 用空格分割字符串，考虑首字母为空格的情况，以及出现任意连续空格的情况
func splitStr(str string) []string {
	l, r, ret := 0, 0, make([]string, 0)
	for r < len(str) {
		if str[r] == ' ' {
			if l != r { // 这是算法核心，能支持首字母为空格，以及出现任意连续空格的情况
				ret = append(ret, str[l:r])
			}
			r++
			l = r
		} else {
			r++
		}
	}

	// TODO 第一版的实现没有考到这点
	// 最后一个没有处理到
	if l != r {
		ret = append(ret, str[l:r])
	}

	return ret
}

// 参考实现，和自己的实现思想一致，只是采用了strings库函数来分割字符串。
func wordPattern_en(pattern string, str string) bool {
	words := strings.Fields(str)
	if len(pattern) != len(words) {
		return false
	}

	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		b := pattern[i]
		s := words[i]
		bs, ok1 := m1[b]
		sb, ok2 := m2[s]
		if !ok1 && !ok2 {
			m1[b] = s
			m2[s] = b
		}

		if ok1 && ok2 {
			if b != sb || s != bs {
				return false
			}
		}

		if (ok1 && !ok2) || (!ok1 && ok2) {
			return false
		}
	}

	return true
}
