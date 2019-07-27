package _15__Kth_Largest_Element_in_an_Array

/*
Runtime: 4 ms, faster than 99.82% of Go online submissions for Kth Largest Element in an Array.
Memory Usage: 3.5 MB, less than 67.50% of Go online submissions for Kth Largest Element in an Array.
*/
// 解法一：采用三路快速排序，标定点固定为数组的第K个元素
func findKthLargest(nums []int, k int) int {
	// 用数组的第K个元素作为分区。
	nums[0], nums[k-1] = nums[k-1], nums[0]
	// 初始[1,l]存放大于val的元素，[l+1,i-1]存放等于val，[i, r-1]待遍历的，【r, n]小于val
	// 采用降序排序
	val, i, l, r := nums[0], 1, 0, len(nums)

	for i < r {
		if nums[i] > val {
			l++
			nums[l], nums[i] = nums[i], nums[l]
			i++
		} else if nums[i] == val {
			i++
		} else {
			r--
			nums[i], nums[r] = nums[r], nums[i]
		}
	}

	nums[0], nums[l] = nums[l], nums[0] // 此时[l,r-1]存放的是等于val的
	if l > k-1 {                        // 左边查找
		return findKthLargest(nums[:l], k)
	} else if r < k { // 右边查找
		return findKthLargest(nums[r:], k-r)
	} else {
		return val
	}
}

// 解法二：最小堆-堆排序，K个元素
