package _44__Reverse_String

/*
Runtime: 664 ms, faster than 42.84% of Go online submissions for Reverse String.
Memory Usage: 8.1 MB, less than 98.82% of Go online submissions for Reverse String.
*/
// 对撞指针
func reverseString(s []byte) {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func reverseString2(s []byte) {
	middle := len(s) >> 1

	for i := 0; i < middle; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
}
