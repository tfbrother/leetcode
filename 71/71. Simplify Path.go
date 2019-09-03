package _1

import "strings"

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Simplify Path.
Memory Usage: 4 MB, less than 100.00% of Go online submissions for Simplify Path.
*/
// 细节：
// 1. 对top进行--时，需要检测top>=0
// 2. paths为空时，期望返回的是"/"，需要特殊处理
func simplifyPath(path string) string {
	paths, top := splitString(path), -1
	//fmt.Println(paths)

	for i := 0; i < len(paths); i++ {
		if paths[i] == ".." {
			if top >= 0 { //
				top--
			}
		} else if paths[i] != "." {
			top++
			paths[top] = paths[i]
		}
	}

	var ret string
	for i := 0; i <= top; i++ {
		ret += "/" + paths[i]
	}

	if ret == "" {
		return "/"
	}
	return ret
}

// 用'/'分割字符串
// 细节：最后处理时，需要检测l和len(path)比较，而不是len(path)-1
func splitString(path string) []string {
	ret, l := make([]string, 0, len(path)>>1), 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if l != i {
				ret = append(ret, string(path[l:i]))
			} // 即有多余的斜杠或者是第一个斜杠
			l = i + 1
		}
	}

	if l != len(path) {
		ret = append(ret, string(path[l:]))
	}

	return ret
}

func simplifyPath_en(path string) string {
	stack := make([]string, 0)

	fields := strings.Split(path, "/")

	for _, f := range fields {
		if f == "." || f == "" {
			continue
		}
		if f == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, f)
		}
	}
	//fmt.Println(stack)
	return "/" + strings.Join(stack, "/")
}
