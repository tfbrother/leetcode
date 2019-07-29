package _25__Valid_Palindrome

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Palindrome.
Memory Usage: 2.7 MB, less than 89.80% of Go online submissions for Valid Palindrome.
*/
// 双索引技术
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		// 首先要保证l，r索引位置的字符是字母和数字字符
		if !isValidChar(s[l]) {
			l++
		} else if !isValidChar(s[r]) {
			r--
		} else if getLowerChar(s[l]) != getLowerChar(s[r]) {
			return false
		} else {
			l++
			r--
		}
	}

	return true
}

func isPalindrome2(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		// 首先要保证l，r索引位置的字符是字母和数字字符
		for l < r && !isValidChar(s[l]) {
			l++
		}
		for l < r && !isValidChar(s[r]) {
			r--
		}

		if getLowerChar(s[l]) != getLowerChar(s[r]) {
			return false
		}

		l++
		r--
	}

	return true
}

// 大写转化成小写
func getLowerChar(c byte) byte {
	if c >= 97 && c <= 122 {
		return c - 32
	}

	return c
}

// 检查是否是有效的字符
func isValidChar(c byte) bool {
	if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) || (c >= 48 && c <= 57) {
		return true
	}

	return false
}
