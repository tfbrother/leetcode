package _02

// https://leetcode-cn.com/problems/happy-number/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Happy Number.
*/
// 核心是如何判断出现循环？利用set来实现
// 解法一：算法思想：参考349
func isHappy1(n int) bool {
	// 余数
	iSet, prod, sum := make(map[int]bool), 0, 0

	for {
		for n > 0 {
			prod, n = n%10, n/10
			sum += prod * prod
		}

		if sum == 1 {
			return true
		} else if iSet[sum] { // 出现循环了
			return false
		} else {
			iSet[sum] = true
			n = sum
			sum = 0 // TODO 这里差点搞忘了写
		}
	}

	// 实际上这里是不可能执行到的吧。
	//return false
}

// 解法一基础上精简代码
func isHappy1_1(n int) bool {
	// 余数
	iSet, prod, sum, ok := make(map[int]bool), 0, 0, false

	for {
		iSet[n], sum = true, 0

		for n > 0 {
			prod, n = n%10, n/10
			sum += prod * prod
		}
		n = sum

		if n == 1 {
			return true
		}
		if _, ok = iSet[n]; ok { // 出现循环了
			return false
		}
	}
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Happy Number.
*/
// 解法二：利用快慢指针思想来判断是否存在环。参考[141]
// O(1)空间复杂度
func isHappy2(n int) bool {
	var digitSquareSum func(num int) int

	digitSquareSum = func(num int) int {
		sum, prod := 0, 0
		for num > 0 {
			prod, num = num%10, num/10
			sum += prod * prod
		}
		return sum
	}

	slow, fast := n, n
	for {
		slow, fast = digitSquareSum(slow), digitSquareSum(digitSquareSum(fast))
		if fast == 1 {
			return true
		}

		if slow == fast {
			return false
		}
	}
}
