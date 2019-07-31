package _09__Minimum_Size_Subarray_Sum

// see https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// 相似题
// 	76. Minimum Window Substring  				Hard
//	325. Maximum Size Subarray Sum Equals k 		Medium
//	718. Maximum Length of Repeated Subarray 	Medium

/*
Runtime: 4 ms, faster than 100.00% of Go online submissions for Minimum Size Subarray Sum.
Memory Usage: 3.9 MB, less than 92.86% of Go online submissions for Minimum Size Subarray Sum.
*/
func minSubArrayLen(s int, nums []int) int {
	// 初始窗口[l...r)
	l, r, sum, min, length := 0, 0, 0, len(nums)+1, 0
	for r < len(nums) {
		sum += nums[r]
		if sum >= s {
			length = r - l + 1
			if min > length {
				min = length
			}
			sum -= nums[l]
			l++
		} else { // < s
			r++
		}
	}

	if min == len(nums)+1 {
		return 0
	}

	return min
}

// 扩展：题意中的大于等于改成等于。即Maximum Size Subarray Sum Equals k
