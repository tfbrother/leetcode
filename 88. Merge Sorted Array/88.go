package _8__Merge_Sorted_Array

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Sorted Array.
Memory Usage: 3.6 MB, less than 81.16% of Go online submissions for Merge Sorted Array.
*/
// 只是归并排序的一个简单扩展，通常归并排序采用的是从前往后，此题考察的是从后往前进行归并
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] < nums2[n-1] {
			nums1[i] = nums2[n-1]
			n--
		} else {
			nums1[i] = nums1[m-1]
			m--
		}
		i--
	}

	for m > 0 {
		nums1[i] = nums1[m-1]
		m--
		i--
	}

	for n > 0 {
		nums1[i] = nums2[n-1]
		n--
		i--
	}
}

// 解法一的基础上优化，仔细思考，为何不用检测 m > 0。
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] < nums2[n-1] {
			nums1[i] = nums2[n-1]
			n--
		} else {
			nums1[i] = nums1[m-1]
			m--
		}
		i--
	}

	for n > 0 {
		nums1[i] = nums2[n-1]
		n--
		i--
	}
}
