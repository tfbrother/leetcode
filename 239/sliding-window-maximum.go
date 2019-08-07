package _39

// see https://leetcode-cn.com/problems/sliding-window-maximum/

/*
Runtime: 852 ms, faster than 19.05% of Go online submissions for Sliding Window Maximum.
Memory Usage: 8.6 MB, less than 100.00% of Go online submissions for Sliding Window Maximum.
*/
// 解法一：暴力算法，每次通过O(k)时间复杂度返回窗口内的最大值
func maxSlidingWindow1(nums []int, k int) []int {
	var ret []int
	if len(nums) == 0 {
		return ret
	}

	for i := 0; i < len(nums)-k+1; i++ {
		max := -1 << 31
		for j := i; j < i+k; j++ {
			if max < nums[j] {
				max = nums[j]
			}
		}

		ret = append(ret, max)
	}

	return ret
}

/*
Runtime: 808 ms, faster than 72.89% of Go online submissions for Sliding Window Maximum.
Memory Usage: 8.5 MB, less than 100.00% of Go online submissions for Sliding Window
*/
// 解法二：该题的核心是求窗口内的最大值，思考如何优化时间复杂度。
// 算法思想：https://mp.weixin.qq.com/s/hiHtzKFamAwoKr3xVRXZMg
// 辅助数据结构：单调队列(基于双向队列实现，双向队列可以基于循环队列和slice实现)
func maxSlidingWindow2(nums []int, k int) []int {
	var ret []int
	if len(nums) == 0 {
		return ret
	}

	moque := &MonotonicQueue{data: make([]int, k)} // 单调队列
	for i := 0; i < len(nums); i++ {
		moque.Push(nums[i])
		if i >= k-1 {
			ret = append(ret, moque.Max())
			moque.Pop(nums[i-k+1])
		}
	}

	return ret
}

// 单调递减队列
type MonotonicQueue struct {
	data  []int
	count int
}

// 向队尾添加元素，会将队列中所有比n小的元素都删掉
// 只会影响到前面添加的元素，后面添加的元素不影响
func (m *MonotonicQueue) Push(n int) {
	// 此处有bug，这种低级错误不应该出现
	// 1. 因为在循环内部有m.count--，所有for语句的条件不应该使用i< m.count。
	// 2. i应该倒序比较
	//for i := 0; i < m.count; i++ {
	//	if m.data[i] < n {
	//		m.count--
	//	}
	//}

	for i := m.count - 1; i >= 0; i-- {
		if m.data[i] < n {
			m.count--
		} else {
			break
		}
	}

	if len(m.data) <= m.count {
		m.data = append(m.data, n)
	} else {
		m.data[m.count] = n
	}
	m.count++
}

// 返回队列中的最大值
func (m *MonotonicQueue) Max() int {
	return m.data[0]
}

// 如果队首元素等于n，就弹出
func (m *MonotonicQueue) Pop(n int) {
	if m.data[0] == n {
		m.data = m.data[1:]
		m.count--
	}
}
