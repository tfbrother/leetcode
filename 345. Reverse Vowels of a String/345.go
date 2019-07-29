package _45__Reverse_Vowels_of_a_String

// 对撞指针，和125题完全是一摸一样的。125是过滤无效字母，此题是过滤掉非元音字母。
// 元音字母：a，e，i，o，u

/**
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Vowels of a String.
Memory Usage: 5.6 MB, less than 37.04% of Go online submissions for Reverse Vowels of a String.
*/
// 解法一：常规思路，
// 	1. 先遍历依次s，把所有元音字母出现的位置记录下来
// 	2. 进行反转
// 	3. 要考虑大小写的情况哦
func reverseVowels(s string) string {
	indexes, length, byteS := make([]int, 0), len(s), []byte(s)
	for i := 0; i < length; i++ {
		// TODO 这段代码如何书写可以更优美呢？
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			indexes = append(indexes, i)
		}
	}

	// 此处也是对撞指针思想
	l, r := 0, len(indexes)-1
	for l < r {
		byteS[indexes[l]], byteS[indexes[r]] = byteS[indexes[r]], byteS[indexes[l]]
		l++
		r--
	}

	return string(byteS)
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Vowels of a String.
Memory Usage: 4.1 MB, less than 100.00% of Go online submissions for Reverse Vowels of a String.
*/
// 对撞指针，在遍历的时候进行处理
// 参考[125]的解法二
func reverseVowels2(s string) string {
	l, r, byteS := 0, len(s)-1, []byte(s)
	for l < r {
		for l < r && !isVowels(byteS[l]) {
			l++
		}

		for l < r && !isVowels(byteS[r]) {
			r--
		}

		if isVowels(byteS[l]) && isVowels(byteS[r]) {
			byteS[l], byteS[r] = byteS[r], byteS[l]
			l++
			r--
		} // TODO else 不做任何处理，思考清楚。
	}

	return string(byteS)
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Vowels of a String.
Memory Usage: 4.1 MB, less than 100.00% of Go online submissions for Reverse Vowels of a String.
*/
// 对撞指针，在遍历的时候进行处理
// 参考[125]的解法二
func reverseVowels3(s string) string {
	l, r, byteS := 0, len(s)-1, []byte(s)
	for l < r {
		if !isVowels(byteS[l]) {
			l++
		} else if !isVowels(byteS[r]) {
			r--
		} else if isVowels(byteS[l]) && isVowels(byteS[r]) {
			byteS[l], byteS[r] = byteS[r], byteS[l]
			l++
			r--
		} // TODO else 不做任何处理，思考清楚。
	}

	return string(byteS)
}

func isVowels(b byte) bool {
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U' {
		return true
	}

	return false
}
