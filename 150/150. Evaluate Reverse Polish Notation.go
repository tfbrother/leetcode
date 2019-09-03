package _50

import (
	"strconv"
)

/*
Runtime: 4 ms, faster than 95.03% of Go online submissions for Evaluate Reverse Polish Notation.
Memory Usage: 4 MB, less than 100.00% of Go online submissions for Evaluate Reverse Polish Notation.
*/
func evalRPN1(tokens []string) int {
	stack, top := make([]int, len(tokens)), -1
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			stack[top-1] = stack[top-1] + stack[top]
			top--
		} else if tokens[i] == "-" {
			stack[top-1] = stack[top-1] - stack[top]
			top--
		} else if tokens[i] == "*" {
			stack[top-1] = stack[top-1] * stack[top]
			top--
		} else if tokens[i] == "/" {
			stack[top-1] = stack[top-1] / stack[top]
			top--
		} else { // 数字
			top++
			stack[top], _ = strconv.Atoi(tokens[i])
		}
	}

	return stack[top]
}

func evalRPN1_1(tokens []string) int {
	stack, top := make([]int, len(tokens)), -1
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			if tokens[i] == "+" {
				stack[top-1] = stack[top-1] + stack[top]
			} else if tokens[i] == "-" {
				stack[top-1] = stack[top-1] - stack[top]
			} else if tokens[i] == "*" {
				stack[top-1] = stack[top-1] * stack[top]
			} else {
				stack[top-1] = stack[top-1] / stack[top]
			}
			top--
		default:
			top++
			stack[top], _ = strconv.Atoi(tokens[i])
		}
	}

	return stack[top]
}

/*
Runtime: 4 ms, faster than 95.03% of Go online submissions for Evaluate Reverse Polish Notation.
Memory Usage: 3.8 MB, less than 100.00% of Go online submissions for Evaluate Reverse Polish Notation.
*/
// 优化空间复杂度
func evalRPN2(tokens []string) int {
	top, num1, num2 := -1, 0, 0
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			num1, _ = strconv.Atoi(tokens[top-1])
			num2, _ = strconv.Atoi(tokens[top])
			tokens[top-1] = strconv.Itoa(num1 + num2)
			top--
		} else if tokens[i] == "-" {
			num1, _ = strconv.Atoi(tokens[top-1])
			num2, _ = strconv.Atoi(tokens[top])
			tokens[top-1] = strconv.Itoa(num1 - num2)
			top--
		} else if tokens[i] == "*" {
			num1, _ = strconv.Atoi(tokens[top-1])
			num2, _ = strconv.Atoi(tokens[top])
			tokens[top-1] = strconv.Itoa(num1 * num2)
			top--
		} else if tokens[i] == "/" {
			num1, _ = strconv.Atoi(tokens[top-1])
			num2, _ = strconv.Atoi(tokens[top])
			tokens[top-1] = strconv.Itoa(num1 / num2)
			top--
		} else { // 数字
			top++
			tokens[top] = tokens[i]
		}
	}

	ret, _ := strconv.Atoi(tokens[top])
	return ret
}

func evalRPN2_1(tokens []string) int {
	top, num1, num2 := -1, 0, 0
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			num1, _ = strconv.Atoi(tokens[top-1])
			num2, _ = strconv.Atoi(tokens[top])
			if tokens[i] == "+" {
				tokens[top-1] = strconv.Itoa(num1 + num2)
			} else if tokens[i] == "-" {
				tokens[top-1] = strconv.Itoa(num1 - num2)
			} else if tokens[i] == "*" {
				tokens[top-1] = strconv.Itoa(num1 * num2)
			} else {
				tokens[top-1] = strconv.Itoa(num1 / num2)
			}
			top--
		default:
			top++
			tokens[top] = tokens[i]
		}
	}

	ret, _ := strconv.Atoi(tokens[top])
	return ret
}
