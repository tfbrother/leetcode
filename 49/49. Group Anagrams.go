package _9

import (
	"sort"
	"strconv"
)

/*
Runtime: 244 ms, faster than 56.37% of Go online submissions for Group Anagrams.
Memory Usage: 340.9 MB, less than 25.00% of Go online submissions for Group Anagrams.
*/
// 查找表+排序
func groupAnagrams(strs []string) [][]string {
	map1, ret := make(map[string][]string, len(strs)>>1), make([][]string, 0, 3)

	for i := 0; i < len(strs); i++ {
		b := bt([]byte(strs[i]))
		sort.Sort(b)
		key := string(b)

		if _, ok := map1[key]; ok {
			map1[key] = append(map1[key], strs[i])
		} else {
			map1[key] = []string{strs[i]}
		}
	}

	for _, item := range map1 {
		ret = append(ret, item)
	}

	return ret
}

type bt []byte

func (s bt) Len() int {
	return len(s)
}

func (s bt) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s bt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 这道题的最优解应该只需要O(N)时间复杂度，需要什么辅助结构呢？
// 而且全部为小写字母，应该可以利用的。
func groupAnagrams1(strs []string) [][]string {
	map1, ret := make(map[string][]string, len(strs)>>1), make([][]string, 0, 3)

	for i := 0; i < len(strs); i++ {
		var b [26]int
		for j := 0; j < len(strs[i]); j++ {
			b[strs[i][j]-'a']++
		}

		var key string
		for k := 0; k < 26; k++ {
			key += strconv.Itoa(b[k])
		}

		if _, ok := map1[key]; ok {
			map1[key] = append(map1[key], strs[i])
		} else {
			map1[key] = []string{strs[i]}
		}
	}

	for _, item := range map1 {
		ret = append(ret, item)
	}

	return ret
}
