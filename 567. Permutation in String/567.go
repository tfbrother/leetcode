package _67__Permutation_in_String

// see https://leetcode-cn.com/problems/permutation-in-string/
// 和[438]解法二一摸一样，注意参数的顺序
// bug：存在matches为负数的情况。所以没有通过。
func checkInclusion1(s2 string, s1 string) bool {
	if len(s1) < len(s2) {
		return false
	}

	l, r, matches, count1, count2 := 0, 0, 0, make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(s2); i++ {
		count2[s2[i]]++
	}

	for r < len(s1) {
		if _, find := count2[s1[r]]; !find {
			r++
			l, count1, matches = r, make(map[byte]int), 0
		} else {
			count1[s1[r]]++
			if count1[s1[r]] == count2[s1[r]] {
				matches++
			}

			r++
			if r-l == len(s2) {
				if len(count2) == matches {
					return true
				}

				if _, find := count2[s1[l]]; find {
					count1[s1[l]]--
					if count1[s1[l]] < count2[s1[l]] {
						matches--
					}
				}

				l++
			}

		}
	}

	return false
}

/*
Runtime: 12 ms, faster than 47.22% of Go online submissions for Permutation in String.
Memory Usage: 4.2 MB, less than 50.00% of Go online submissions for Permutation in String.
*/
// fix bug
func checkInclusion1_1(s2 string, s1 string) bool {
	if len(s1) < len(s2) {
		return false
	}

	l, r, matches, count1, count2 := 0, 0, 0, make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(s2); i++ {
		count2[s2[i]]++
	}

	for r < len(s1) {
		if _, find := count2[s1[r]]; !find {
			r++
			l, count1, matches = r, make(map[byte]int), 0
		} else {
			count1[s1[r]]++
			if count1[s1[r]] == count2[s1[r]] {
				matches++
			}

			r++
			if r-l == len(s2) {
				if len(count2) == matches {
					return true
				}

				if count1[s1[l]] == count2[s1[l]] {
					matches--
				}

				count1[s1[l]]--
				l++
			}

		}
	}

	return false
}

/*
Runtime: 12 ms, faster than 47.22% of Go online submissions for Permutation in String.
Memory Usage: 4.2 MB, less than 50.00% of Go online submissions for Permutation in String.
*/
// 和[438]解法四一摸一样
func checkInclusion2(s2 string, s1 string) bool {
	if len(s1) < len(s2) {
		return false
	}

	l, r, matches, count1, count2 := 0, 0, 0, make(map[byte]int), make(map[byte]int)

	for i := 0; i < len(s2); i++ {
		count2[s2[i]]++
	}

	for r < len(s1) {
		if _, find := count2[s1[r]]; !find {
			r++
			l, count1, matches = r, make(map[byte]int), 0
		} else {
			count1[s1[r]]++
			if count1[s1[r]] == count2[s1[r]] {
				matches++
			}

			r++
			for len(count2) == matches {
				if r-l == len(s2) {
					return true
				}

				if _, find := count2[s1[l]]; find {
					count1[s1[l]]--
					if count1[s1[l]] < count2[s1[l]] {
						matches--
					}
				}

				l++
			}

		}
	}

	return false
}

// 参考自英文版0ms的答案
// 此解法很精妙，438也可以采用
func checkInclusion_en(s1 string, s2 string) bool {
	windowWidth, n := len(s1), len(s2)
	if n < windowWidth {
		return false
	}

	// prepare the pattern and the initial sliding window
	pattern := [26]int{}
	window := [26]int{}
	for i := 0; i < windowWidth; i++ {
		pattern[int(s1[i]-'a')]++
		window[int(s2[i]-'a')]++
	}

	for l := 0; l < n-windowWidth+1; l++ {
		if l > 0 {
			window[int(s2[l-1]-'a')]--
			window[int(s2[l+windowWidth-1]-'a')]++ // r = l+windowWidth-1
		}
		// TODO 直接利用golang的数组比较，比较耗时
		if window == pattern {
			return true
		}
	}
	return false
}

// 参考自国内区0ms的答案
func checkInclusion_cn(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	h1, h2 := 0, 0
	for i := 0; i < n1; i++ {
		c1 := s1[i] - 'a'
		c2 := s2[i] - 'a'
		// 位运算
		h1 += 1 << c1
		h2 += 1 << c2
	}

	if h1 == h2 {
		return true
	}

	for i := n1; i < n2; i++ {
		cb := s2[i-n1] - 'a'
		ce := s2[i] - 'a'
		// 利用 cb 和 ce 更新 h2
		h2 += (1 << ce) - (1 << cb)
		if h1 == h2 {
			return true
		}
	}

	return false
}
