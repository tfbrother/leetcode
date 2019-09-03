package _0

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Valid Parentheses.
*/
func isValid(s string) bool {
	stack, top := make([]byte, len(s)), -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			top++
			stack[top] = s[i]
		} else {
			if top < 0 {
				return false
			}

			var match byte
			if s[i] == ']' {
				match = '['
			} else if s[i] == ')' {
				match = '('
			} else if s[i] == '}' {
				match = '{'
			} else {
				return false
			}

			if stack[top] != match {
				return false
			}

			top-- // 弹出栈
		}
	}

	if top >= 0 {
		return false
	}

	return true
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Valid Parentheses.
*/
func isValid2(s string) bool {
	arr := make([]string, len(s))
	l := 0

	for _, v := range s {
		a := string(v)
		if a == "(" || a == "[" || a == "{" {
			arr[l] = a //放入栈
			l++
		} else {
			if l == 0 {
				return false
			}
			b := arr[l-1] //取栈顶的元素
			if (a == ")" && b != "(") || (a == "}" && b != "{") || (a == "]" && b != "[") {
				return false
			}
			// 弹出栈顶元素
			l--
		}
	}

	return l == 0
}
