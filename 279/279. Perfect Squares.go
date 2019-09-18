package _79

import "math"

/*
Runtime: 8 ms, faster than 95.15% of Go online submissions for Perfect Squares.
Memory Usage: 6.4 MB, less than 33.33% of Go online submissions for Perfect Squares.
*/
// 解法一：图论，求最短路径
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func numSquares1(n int) int {
	queue, visted := make([][2]int, 1), make([]bool, n+1)
	queue[0] = [2]int{n, 0}

	for len(queue) > 0 {
		num, step := queue[0][0], queue[0][1]
		queue = queue[1:]

		for i := 1; num-i*i >= 0; i++ {
			v := num - i*i
			if !visted[v] {
				if v == 0 {
					return step + 1
				}
				queue = append(queue, [2]int{v, step + 1})
				visted[v] = true
			}
		}
	}

	return -1
}

/*
Runtime: 20 ms, faster than 85.46% of Go online submissions for Perfect Squares.
Memory Usage: 5.7 MB, less than 100.00% of Go online submissions for Perfect Squares.
*/
// 解法二：动态规划
func numSquares2(n int) int {
	// 递推公式
	// nums[n] = min(1+nums[n-i^2])

	nums := make([]int, n+1)
	nums[1] = 1

	for i := 2; i <= n; i++ {
		minNum := int(math.MaxInt32)
		for j := 1; j*j <= i; j++ {
			minNum = min(minNum, 1+nums[i-j*j])
		}

		nums[i] = minNum
	}

	return nums[n]
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
