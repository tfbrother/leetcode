package _38

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
func FindAnagrams1_4(s string, p string) []int {
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

/*
Runtime: 108 ms, faster than 48.60% of Go online submissions for Find All Anagrams in a String.
Memory Usage: 7.9 MB, less than 92.86% of Go online submissions for Find All Anagrams in a String.
*/
// 第五版，在四的基础上精简代码，但没有利用题目中的条件：字符串只包含小写英文字母。
func FindAnagrams1_5(s string, p string) []int {
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

/*
Runtime: 100 ms, faster than 58.43% of Go online submissions for Find All Anagrams in a String.
Memory Usage: 8 MB, less than 91.67% of Go online submissions for Find All Anagrams in a String.
*/
// 解法二：主要的优化在于判断当前窗口是否满足条件，引入了一个新的变量matches
// TODO 核心是思考清楚，matches和窗口的长度，是怎能能判断当前窗口是满足条件的。同时matches的维护也是核心，各种情况要考虑清楚
// 虽然通过了，但其实该题是有bug的。用该题的解法去解[567]的时候没有通过。具体查看567的第一次提交。
// bug：存在matches为负数的情况。仔细思考啊。
// 算法思想：
//	1. 循环中始终保证固定窗口大小。
//	2. 检测窗口内的元素是否满足条件，此处使用的是matches变量来记录。
//  3. matches满足条件，则加入结果集。将窗口往右边移动一位。移动的时候要维护matches变量。
//	4. 优化：加速窗口的移动。比如遇到位置r的元素不在p中，则立马将窗口起始位置移动到r+1去。
func FindAnagrams2(s string, p string) []int {
	if len(s) < len(p) {
		return make([]int, 0)
	}

	counts, countp, i, l, r, ret, matches := make(map[byte]int), make(map[byte]int), 0, 0, 0, make([]int, 0), 0
	for ; i < len(p); i++ {
		countp[p[i]]++
	}

	for r < len(s) {
		if _, find := countp[s[r]]; !find { // s[r]不在p中，移动窗口，提高了效率，是窗口移动速度更快，但是增加了内存使用。
			r++
			l, counts, matches = r, make(map[byte]int), 0
		} else { // s[r]在p中
			counts[s[r]]++
			if counts[s[r]] == countp[s[r]] {
				matches++
			}

			r++

			// 检查当前窗口的大小是否满足题意
			if r-l == len(p) { // (1)
				// 检测当前窗口的内容是否满足题意，此时利用matches来判断
				if matches == len(countp) { // (2) TODO (1),(2)两个条件的顺序是否会影响解答？
					ret = append(ret, l) // 此时s[l]肯定在p中，需要matches--，留在下面一并处理。
				}

				// 此时s[l]不一定在p中，当前窗口：[b,b,c],counts：[a,b,c]，这种情况就不需要matches--
				if _, find := countp[s[l]]; find {

					/*
						// bug：存在matches为负数的情况。
						// 假设初始counts[s[l]]=3, countp[s[l]]=2，那么这种情况matches是不能减的
						counts[s[l]]--

						if countp[s[l]] > counts[s[l]] {
							matches--
						}
					*/

					// fix bug
					if countp[s[l]] == counts[s[l]] {
						matches--
					}
					counts[s[l]]--
				}

				l++ // 移动窗口，维持窗口的最大值不能超过len(p)
			}
		}
	}

	return ret
}

// 解法三：参考自其它答案
// 优化内存使用，因为题目中字符串只包含小写英文字母，这个条件还没用到。
// 	1. 首先matches的含义不一样了。
// TODO 算法思想尚未理解清楚，该解法目前是性能最高的解法。
func FindAnagrams3(s string, p string) []int {
	if len(s) < len(p) {
		return make([]int, 0)
	}

	var count [26]int // 前提：字符串只包含小写英文字母
	// matches为窗口中匹配的字母数量
	matches, result := 0, make([]int, 0, 5)
	for i := 0; i < len(p); i++ {
		count[p[i]-'a']++
	}

	for l, r := 0, 0; r < len(s); r++ {
		count[s[r]-'a']-- // 可能为负数哦
		// 初始情况下count中大于0的表示p中的字符，如果对其进行减1后
		// 其值大于等于0，则说明s[r]在p中。
		if (count[s[r]-'a']) >= 0 {
			matches++
		}

		// 1. 维持窗口的大小不变
		// 2. 窗口中做记录：matches的定义是什么？窗口中匹配字母的个数
		if r >= len(p) {
			// 如果s[l]不在p中，则会先count[s[l]-'a']--，后面才会对count[s[l]-'a']++
			// 此时count[s[l]-'a']等于0。如果count[s[l]-'a']>0，则说明s[l]在p中
			// 所以在窗口的移动过程中，要维持matches变量的定义不变。
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

// 别人的答案
// 算法思想：
// 	1.先找到一个窗口，窗口满足第一个条件：即p中的字母以及个数都包含在窗口内。(实现时用的是match)
//	2.然后不断缩小窗口，即i++，但是必须确保第一个条件始终满足。
// 	3.当满足窗口缩小长度等于len(p)时仍然满足match == len(needs)，则加入结果。
// 	4.否则继续寻找满足第一个条件窗口
func FindAnagrams4(s string, p string) []int {
	var (
		result      []int = []int{}
		left, right int   = 0, 0
		match       int   = 0
	)
	needs := make(map[byte]int)
	windows := make(map[byte]int)
	for _, val := range p {
		needs[byte(val)]++
	}

	for right < len(s) {
		c1 := s[right]
		if _, ok := needs[c1]; ok {
			windows[c1]++
			if windows[c1] == needs[c1] {
				match++
			}
		}
		right++

		for match == len(needs) {
			if right-left == len(p) {
				result = append(result, left)
			}

			c2 := s[left]
			if _, ok := needs[c2]; ok {
				windows[c2]--
				if windows[c2] < needs[c2] {
					match--
				}
			}
			left++
		}
	}
	return result
}

/*
Runtime: 80 ms, faster than 98.88% of Go online submissions for Find All Anagrams in a String.
Memory Usage: 7.9 MB, less than 91.67% of Go online submissions for Find All Anagrams in a
*/
// 参考[567]英文版0ms的答案，checkInclusion_en，
func FindAnagrams5(s string, p string) []int {
	n, winLen, ret := len(s), len(p), make([]int, 0)
	if n < winLen {
		return ret
	}

	var (
		window, pattern [26]int
		l, r, i         int
	)

	for ; i < winLen; i++ {
		window[int(s[i]-'a')]++
		pattern[int(p[i]-'a')]++
	}

	for l < n-winLen+1 {
		r = l + winLen - 1
		if l > 0 {
			window[s[l-1]-'a']--
			window[s[r]-'a']++
		}

		// TODO 直接利用golang的数组比较，这一步会比较耗时，可以优化。
		if window == pattern {
			ret = append(ret, l)
		}
		l++
	}

	return ret
}

/*
Runtime: 80 ms, faster than 98.88% of Go online submissions for Find All Anagrams in a String.
Memory Usage: 8 MB, less than 91.67% of Go online submissions for Find All Anagrams in a String.
*/
// TODO 虽然ac了，但是存在bug。
// 解法六：参考[567]国内版0ms的答案，checkInclusion_cn，
// 主要优化就是在解法五的基础上优化窗口比较，利用位运算来优化。比解法五提升约6倍的速度。不过位运算存在溢出的情况。
//  两个重点：
//		1. 可以利用位运算计算出的值作为判断条件吗？
//			其实不可以，比如：s, p = "bbb", "caa" 此时计算出来: h1 == h2 == 6
//		2. 溢出了如何处理？
// 			'z'-'a'=25，对于int32来说，128(2^7)个z就会导致溢出。int64的话需要2^39个z才会导致溢出，因此对于int64可以不考虑溢出。
// 	BenchmarkFindAnagrams51e2-4   	    2000	    686405 ns/op
//	BenchmarkFindAnagrams61e2-4   	   20000	    105515 ns/op
func FindAnagrams6(s string, p string) []int {
	n, winLen, ret := len(s), len(p), make([]int, 0)
	if n < winLen {
		return ret
	}

	var (
		h1, h2 int
		r, i   int
	)

	for ; i < winLen; i++ {
		h1 += 1 << (s[i] - 'a')
		h2 += 1 << (p[i] - 'a')
	}

	if h1 == h2 {
		ret = append(ret, 0)
	}

	for r = winLen; r < n; r++ {
		h1 += 1<<(s[r]-'a') - 1<<(s[r-winLen]-'a')

		if h1 == h2 {
			ret = append(ret, r-winLen+1)
		}
	}

	return ret
}

/*
BenchmarkFindAnagrams1_41e2-4   	     200	   6399568 ns/op
BenchmarkFindAnagrams1_51e2-4   	     200	   6614262 ns/op
BenchmarkFindAnagrams21e2-4     	     200	   8354550 ns/op
BenchmarkFindAnagrams31e2-4     	   10000	    146473 ns/op
BenchmarkFindAnagrams41e2-4     	     200	   8130003 ns/op
BenchmarkFindAnagrams51e2-4     	    2000	    665992 ns/op
BenchmarkFindAnagrams61e2-4     	   20000	     98353 ns/op
*/
