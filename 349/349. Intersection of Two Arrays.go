package _49

// see https://leetcode-cn.com/problems/intersection-of-two-arrays/

/*
Runtime: 4 ms, faster than 86.02% of Go online submissions for Intersection of Two Arrays.
Memory Usage: 4.4 MB, less than 100.00% of Go online submissions for Intersection of Two Arrays.
*/
// 解法一：set解决，只是go没有提供set，所以把map当set使用
func intersection(nums1 []int, nums2 []int) []int {
	map1, ret, i := make(map[int]bool), make([]int, 0), 0

	for i = 0; i < len(nums1); i++ {
		map1[nums1[i]] = true
	}

	for i = 0; i < len(nums2); i++ {
		if map1[nums2[i]] {
			ret = append(ret, nums2[i])
			map1[nums2[i]] = false
		}
	}

	return ret
}
