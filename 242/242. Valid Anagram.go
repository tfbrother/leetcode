package _42

// https://leetcode-cn.com/problems/valid-anagram/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
Memory Usage: 3 MB, less than 83.33% of Go online submissions for Valid Anagram.
*/
// 解法一：采用map来解决，解法参考350
func isAnagram1(s string, t string) bool {
	var (
		freq [26]int
		i    int
	)

	for i = 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i = 0; i < len(t); i++ {
		freq[t[i]-'a']--
		if freq[t[i]-'a'] < 0 {
			return false
		}
	}

	for i = 0; i < 26; i++ {
		if freq[i] > 0 {
			return false
		}
	}

	return true
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
Memory Usage: 3 MB, less than 100.00% of Go online submissions for Valid Anagram.
*/
// 在解法一的基础上精简代码
func isAnagram1_1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var (
		freq [26]int
		i    int
	)

	for i = 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	for i = 0; i < len(t); i++ {
		freq[t[i]-'a']--
		if freq[t[i]-'a'] < 0 {
			return false
		}
	}

	return true
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
Memory Usage: 3 MB, less than 100.00% of Go online submissions for Valid Anagram.
*/
func isAnagram2(s string, t string) bool {
	l1, l2 := len(s), len(t)
	if l1 != l2 {
		return false
	}

	var freq1, freq2 [26]int
	for i := 0; i < l1; i++ {
		freq1[s[i]-'a']++
	}
	for j := 0; j < l2; j++ {
		freq2[t[j]-'a']++
	}

	for k := 0; k < 26; k++ {
		if freq1[k] != freq2[k] {
			return false
		}
	}

	return true
}
