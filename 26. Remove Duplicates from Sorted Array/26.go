package _6__Remove_Duplicates_from_Sorted_Array

/*
Runtime: 44 ms, faster than 88.08% of Go online submissions for Remove Duplicates from Sorted Array.
Memory Usage: 7.5 MB, less than 88.66% of Go online submissions for Remove Duplicates from Sorted Array.
*/
// 和27题基本一样
// 扩展题目：83，82。处理的是有序链表
func removeDuplicates(nums []int) int {
	// 就是当nums为空时或者只有一个值的时候，直接返回。
	if len(nums) <= 1 {
		return len(nums)
	}

	k, n := 0, len(nums) // [0, k]为非重复元素
	for i := 1; i < n; i++ {
		if nums[i] != nums[k] {
			k++
			// 针对全部元素都不重复的情况进行优化
			if k != i {
				nums[k] = nums[i]
			}
		}
	}

	return k + 1
}
