package _63

// https://leetcode-cn.com/problems/ugly-number/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Ugly Number.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Ugly Number.
*/
func isUgly(n int) bool {
	if n == 0 { // 第一版超时就是因为没有对0特殊处理
		return false
	}
	for {
		for n%2 == 0 {
			n /= 2
		}

		for n%3 == 0 {
			n /= 3
		}

		for n%5 == 0 {
			n /= 5
		}

		if n == 1 {
			return true
		} else {
			return false
		}
	}
}
