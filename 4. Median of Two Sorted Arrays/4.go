package main

import "fmt"

// 直观的想法就是归并排序，两个数据是有序的，但是时间复杂度是O(m + n)，不满足题意。
// 解法一：时间空间都是O(m + n)
/*
Runtime: 32 ms, faster than 8.09% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5.9 MB, less than 35.23% of Go online submissions for Median of Two Sorted Arrays.
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		middle := len(nums2) >> 1
		if len(nums2)%2 == 0 { // 偶数
			return float64(nums2[middle]+nums2[middle-1]) / 2.0
		} else {
			return float64(nums2[middle])
		}
	} else if len(nums2) == 0 {
		middle := len(nums1) >> 1
		if len(nums1)%2 == 0 { // 偶数
			return float64(nums1[middle]+nums1[middle-1]) / 2.0
		} else {
			return float64(nums1[middle])
		}
	}

	// nums1, nums2 都不为空
	tmp, l, r, i := make([]int, len(nums1)+len(nums2)), 0, 0, 0
	for l < len(nums1) && r < len(nums2) {
		if nums1[l] <= nums2[r] {
			tmp[i] = nums1[l]
			l++
		} else {
			tmp[i] = nums2[r]
			r++
		}

		i++
	}

	for l < len(nums1) {
		tmp[i] = nums1[l]
		i++
		l++
	}

	for r < len(nums2) {
		tmp[i] = nums2[r]
		i++
		r++
	}
	fmt.Println(tmp)
	middle := len(tmp) >> 1
	if len(tmp)%2 == 0 { // 偶数
		return float64(tmp[middle]+tmp[middle-1]) / 2.0
	} else {
		return float64(tmp[middle])
	}
}
