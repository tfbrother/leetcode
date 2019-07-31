package _38__Find_All_Anagrams_in_a_String

// see https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
// 相似题
// 	242. Valid Anagram 				Easy
//	567. Permutation in String		Medium

// 滑动窗口
// 细节：注意p中可能包含重复字符哦。
// 异位词的定义：
//	用例："cbaebabacdabc"，"abc" 期望输出：[0,6,10]，可见只要字符集合相同即返回，所以abc也满足条件，不要真正的异位词。
// 第一版实现返回[]，预期[0,6,10]
func findAnagrams1_1(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	// i,j为滑动窗口的左右边界，width表示窗口的最大值。
	i, j, width, charCount1, charCount2, ret := 0, 0, len(p)-1, make(map[byte]int), make(map[byte]int), make([]int, 0)
	for k := 0; k <= width; k++ {
		charCount2[p[k]]++
	}

	find, is := false, true
	for j < len(s) && j <= i+width {
		if _, find = charCount2[s[j]]; !find {
			if j-i == width { // 找到一个窗口，检测是否满足条件
				for char, count1 := range charCount1 {
					if charCount2[char] != count1 { // 非异位词
						is = false
						break
					}
				}

				if is {
					ret = append(ret, i)
				}
			}

			j++
			i = j + 1
		} else { // 查找到了
			charCount1[s[j]]++
			j++ // 第一版实现超出时间限制，就是忘了这句，太大意了。
		}
	}

	return ret
}

// 第二版本实现返回[0]
func findAnagrams1_2(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	// i,j为滑动窗口的左右边界，width表示窗口的最大值。
	i, j, width, charCount1, charCount2, ret := 0, 0, len(p)-1, make(map[byte]int), make(map[byte]int), make([]int, 0)
	for k := 0; k <= width; k++ {
		charCount2[p[k]]++
	}

	find, is := false, true
	for j < len(s) && j <= i+width {
		if _, find = charCount2[s[j]]; !find {
			j++
			i = j + 1
			charCount1 = make(map[byte]int)
		} else { // 查找到了
			charCount1[s[j]]++
			j++               // 第一版实现超出时间限制，就是忘了这句，太大意了。
			if j-i == width { // 找到一个窗口，检测是否满足条件
				for char, count1 := range charCount1 {
					if charCount2[char] != count1 { // 非异位词
						is = false
						break
					}
				}
				charCount1 = make(map[byte]int)
				if is {
					ret = append(ret, i)
				}
			}
		}
	}

	return ret
}

// 第三版本实现返回
// 归根结底是要思考清楚滑动窗口如何滑动
// 在[i...i+width]之间如何移动？
// 	1.当s[j]不再p中时，j++，i=j+1，清空charCount1
//  2.当s[j]在p中时，分两种情况
// 		a.如果j不等于i+width，则继续j++
//		b.j == i+width，需要比较窗口内的所有元素出现次数是否与p相同。相同则加入结果集，不相同则如何处理呢？这是核心。
//		  不相同，则且charCount1[s[i]]--，i++，j++

//  第三版本实现返回[0]
func findAnagrams1_3(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	// i,j为滑动窗口的左右边界，width表示窗口的最大值。
	i, j, width, charCount1, charCount2, ret := 0, 0, len(p)-1, make(map[byte]int), make(map[byte]int), make([]int, 0)
	for k := 0; k <= width; k++ {
		charCount2[p[k]]++
	}

	find, is := false, true
	for j < len(s) && j <= i+width {
		// fmt.Println("i===", i, ";j===", j)
		if _, find = charCount2[s[j]]; !find {
			j++
			i = j + 1
			charCount1 = make(map[byte]int)
		} else { // 查找到了
			charCount1[s[j]]++
			// 第一版实现超出时间限制，就是忘了这句，太大意了。
			if j == i+width { // 找到一个窗口，检测是否满足条件
				for char, count := range charCount1 { // 第三版用的是charCount2，这版用的是charCount1
					if charCount2[char] != count { // 非异位词
						is = false
						break
					}
				}

				if is {
					ret = append(ret, i)
					is = false
				}
				charCount1[s[i]]--
				i++
			}
			j++
		}
	}

	return ret
}

// 第四版通过
func findAnagrams1_4(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	// i,j为滑动窗口的左右边界，width表示窗口的最大值。
	i, j, width, charCount1, charCount2, ret := 0, 0, len(p)-1, make(map[byte]int), make(map[byte]int), make([]int, 0)
	for k := 0; k <= width; k++ {
		charCount2[p[k]]++
	}

	find, is := false, true
	for j < len(s) && j <= i+width {
		// fmt.Println("i===", i, ";j===", j)
		if _, find = charCount2[s[j]]; !find { // 窗口从j+1开始
			j++
			i = j // 第三版这里错了
			charCount1 = make(map[byte]int)
		} else { // 查找到了
			charCount1[s[j]]++
			if j == i+width { // 找到一个窗口，检测是否满足条件
				for char, count1 := range charCount1 {
					if charCount2[char] != count1 { // 这里有优化空间，可以将窗口移动到char索引位置后面一个去。
						is = false
						break
					}
				}

				if is {
					ret = append(ret, i)
				}
				is = true // 第三版这里错了
				// 窗口右移动一位
				charCount1[s[i]]--
				i++
			}
			j++ // 第一版实现超出时间限制，就是忘了这句，太大意了。
		}
	}

	return ret
}

// 第五版，在四的基础上精简代码，但没有利用题目中的条件：字符串只包含小写英文字母。
func findAnagrams1_5(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	// i,j为滑动窗口的左右边界，width表示窗口的最大值。
	i, j, width, charCount1, charCount2, ret := 0, 0, len(p)-1, make(map[byte]int), make(map[byte]int), make([]int, 0)
	for k := 0; k <= width; k++ {
		charCount2[p[k]]++
	}

	for j < len(s) && j <= i+width {
		// fmt.Println("i===", i, ";j===", j)
		if _, find := charCount2[s[j]]; !find { // 窗口从j+1开始
			j++
			i = j // 第三版这里错了
			charCount1 = make(map[byte]int)
		} else { // 查找到了
			charCount1[s[j]]++

			if j == i+width { // 找到一个窗口，检测是否满足条件
				if isEqual(charCount1, charCount2) {
					ret = append(ret, i)
				}
				// 窗口右移动一位
				charCount1[s[i]]--
				i++
			}
			j++ // 第一版实现超出时间限制，就是忘了这句，太大意了。
		}
	}

	return ret
}

func isEqual(a, b map[byte]int) bool {
	for k, v := range b {
		if a[k] != v {
			return false
		}
	}
	return true
}

// 参考自其它答案
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return make([]int, 0)
	}

	var count [26]int // 前提：字符串只包含小写英文字母
	// matches为窗口中匹配的字母数量
	matches, result := 0, make([]int, 0, 5)
	for i := 0; i < len(p); i++ {
		count[p[i]-'a']++
	}

	l := 0
	for r := 0; r < len(s); r++ {
		count[s[r]-'a']-- // 可能为负数哦
		if (count[s[r]-'a']) >= 0 {
			matches++
		}

		if r >= len(p) {
			count[s[l]-'a']++
			if (count[s[l]-'a']) > 0 {
				matches--
			}
			l++
		}

		if matches == len(p) {
			result = append(result, l)
		}
	}

	return result
}
