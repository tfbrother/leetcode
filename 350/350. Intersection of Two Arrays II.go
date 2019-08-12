package _50

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Intersection of Two Arrays II.
Memory Usage: 4.8 MB, less than 100.00% of Go online submissions for Intersection of Two Arrays II.
*/
// 解法一：使用map
func intersect(nums1 []int, nums2 []int) []int {
	map1, ret, i, count, ok := make(map[int]int), make([]int, 0), 0, 0, false

	for i = 0; i < len(nums1); i++ {
		map1[nums1[i]]++
	}

	for i = 0; i < len(nums2); i++ {
		if count, ok = map1[nums2[i]]; ok && count > 0 {
			ret = append(ret, nums2[i])
			map1[nums2[i]]--
		}
	}

	return ret
}
