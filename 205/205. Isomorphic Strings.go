package _05

// https://leetcode-cn.com/problems/isomorphic-strings

/*
Runtime: 4 ms, faster than 56.32% of Go online submissions for Isomorphic Strings.
Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Isomorphic Strings.
*/
// 和290一摸一样
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS, mapT := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		si, ok1 := mapS[s[i]]
		ti, ok2 := mapT[t[i]]

		if !ok1 && !ok2 {
			mapS[s[i]], mapT[t[i]] = t[i], s[i]
		} else if si != t[i] || ti != s[i] {
			return false
		}
	}

	return true
}
