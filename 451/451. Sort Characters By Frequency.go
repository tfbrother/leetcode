package _51

import "sort"

// https://leetcode-cn.com/problems/sort-characters-by-frequency/

/*
Runtime: 460 ms, faster than 22.22% of Go online submissions for Sort Characters By Frequency.
Memory Usage: 8.1 MB, less than 100.00% of Go online submissions for Sort Characters By Frequency.
*/
// 核心是map的排序，go的map底层是hash，是不支持排序的。
// 所以需要多一层映射，将map转化成slice。
// freqs的设计是核心，因为sort排序会打乱索引，所以freqs还需要保存key的值
func FrequencySort1(s string) string {
	charMap, freqs, ret := make(map[byte]int), make([][2]int, 0), ""
	for i := 0; i < len(s); i++ {
		if c, ok := charMap[s[i]]; ok {
			freqs[c][1]++
		} else { // 第一次出现
			charMap[s[i]] = len(freqs)
			freqs = append(freqs, [2]int{int(s[i]), 1}) // 第一次出现
		}
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i][1] > freqs[j][1]
	})

	for i := 0; i < len(freqs); i++ {
		c := string(byte(freqs[i][0]))
		for j := 0; j < freqs[i][1]; j++ {
			ret += c
		}
	}

	return ret
}

/*
Runtime: 244 ms, faster than 31.31% of Go online submissions for Sort Characters By Frequency.
Memory Usage: 7.8 MB, less than 100.00% of Go online submissions for Sort Characters By Frequency.
*/
// 解法一基础上性能优化，主要是针对freqs
func FrequencySort1_1(s string) string {
	charMap, freqs, ret := make(map[byte]int), make([]string, 0), ""
	for i := 0; i < len(s); i++ {
		if c, ok := charMap[s[i]]; ok {
			freqs[c] += string(s[i])
		} else {
			charMap[s[i]] = len(freqs)
			freqs = append(freqs, string(s[i])) // 出现一次
		}
	}

	sort.Slice(freqs, func(i, j int) bool {
		return len(freqs[i]) > len(freqs[j])
	})

	for i := 0; i < len(freqs); i++ {
		ret += freqs[i]
	}

	return ret
}

/*
Runtime: 8 ms, faster than 73.74% of Go online submissions for Sort Characters By Frequency.
Memory Usage: 6.1 MB, less than 100.00% of Go online submissions for Sort Characters By Frequency.
*/
// 解法一基础上性能优化，分别求解charMap和freqs
// 同时减少使用字符串的拼接
func FrequencySort1_2(s string) string {
	charMap, ret := make(map[rune]int), ""
	for _, v := range s {
		charMap[v]++
	}

	freqs := make([]string, len(charMap))
	for k, v := range charMap {
		tmp := make([]rune, v)
		for i := range tmp {
			tmp[i] = k
		}
		freqs = append(freqs, string(tmp))
	}

	sort.Slice(freqs, func(i, j int) bool {
		return len(freqs[i]) > len(freqs[j])
	})

	for i := 0; i < len(freqs); i++ {
		ret += freqs[i]
	}

	return ret
}

/*
Runtime: 4 ms, faster than 95.96% of Go online submissions for Sort Characters By Frequency.
Memory Usage: 4.9 MB, less than 100.00% of Go online submissions for Sort Characters By
*/
// 解法一基础上性能优化，分别求解charMap和freqs
// 优化到不使用字符串的拼接
func FrequencySort1_3(s string) string {
	charMap, freqs := make(map[byte]int), make([][2]int, 0)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}

	for k, v := range charMap {
		freqs = append(freqs, [2]int{int(k), v})
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i][1] > freqs[j][1]
	})

	ret, k := make([]byte, len(s)), 0
	for i := 0; i < len(freqs); i++ {
		c := byte(freqs[i][0])
		for j := 0; j < freqs[i][1]; j++ {
			ret[k+j] = c
		}
		k += freqs[i][1]
	}

	return string(ret)
}

/*
Runtime: 4 ms, faster than 95.96% of Go online submissions for Sort Characters By Frequency.
Memory Usage: 4.9 MB, less than 100.00% of Go online submissions for Sort Characters By Frequency.
*/
// 解法1_3基础上性能优化，
// 减少频繁的动态申请内存
// TODO BEST VERSION 该版本比frequencySort_en版本的时间和内存占用都低。
func FrequencySort1_4(s string) string {
	charMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}

	freqs, i := make([][2]int, len(charMap), len(charMap)), 0
	for k, v := range charMap {
		freqs[i] = [2]int{int(k), v}
		i++
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i][1] > freqs[j][1]
	})

	ret, k := make([]byte, len(s)), 0
	for i := 0; i < len(freqs); i++ {
		c := byte(freqs[i][0])
		for j := 0; j < freqs[i][1]; j++ {
			ret[k+j] = c
		}
		k += freqs[i][1]
	}

	return string(ret)
}

/*
Runtime: 8 ms, faster than 73.74% of Go online submissions for Sort Characters By Frequency.
Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Sort Characters By Frequency.
*/
// TODO 来自参考答案，思路和解法一一摸一样，为何性能差别那么大？
func FrequencySort_en(s string) string {
	fre := make(map[rune]int, len(s))
	for _, v := range s {
		fre[v]++
	}

	ss := make([]string, 0, len(fre))
	for k, v := range fre {
		tmp := make([]rune, v)
		for i := range tmp {
			tmp[i] = k
		}
		ss = append(ss, string(tmp))
	}

	sort.Sort(str(ss))

	res := ""
	for _, v := range ss {
		res += v
	}

	return res

}

type str []string

func (s str) Len() int {
	return len(s)
}

func (s str) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func (s str) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
