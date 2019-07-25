package _0__Remove_Duplicates_from_Sorted_Array_II

/*
Runtime: 4 ms, faster than 94.66% of Go online submissions for Remove Duplicates from Sorted Array II.
Memory Usage: 6.1 MB, less than 53.49% of Go online submissions for Remove Duplicates from Sorted Array II.
*/
// 和26题相似，支持使得每个元素最多出现任意次。
// 算法思想：
// 	设置两个变量，k和count。k定义为满足条件的最后一个元素索引位置，即[0,k]保存所以满足条件的元素，
// 	count记录nums[k]在nums[0：k+1]中出现的次数。这里的count最大值为X，就可以支持每个元素最多出现X次，进行扩展。
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	// [0,k]保存所以满足条件的元素，count记录nums[k]在nums[0,k]中出现的次数。
	// 这里的count最大值为X，就可以支持每个元素最多出现X次，进行扩展。
	k, count, x := 0, 1, 2

	for i := 1; i < n; i++ {
		if nums[i] == nums[k] {
			if count < x { // nums[i] 需要包含进来
				k++
				count++
				nums[k] = nums[i]
			}
			// count == x 时，nums[i]丢弃
		} else {
			k++
			count = 1
			nums[k] = nums[i]
		}
	}

	return k + 1
}

/*
Runtime: 4 ms, faster than 94.66% of Go online submissions for Remove Duplicates from Sorted Array II.
Memory Usage: 6.1 MB, less than 53.49% of Go online submissions for Remove Duplicates from Sorted Array II.
*/
// 因为最多出现两次，所以可以枚举。不用count变量
func removeDuplicates2(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	// [0,k]保存所以满足条件的元素
	k := 1
	for i := 2; i < n; i++ {
		// if nums[i] != nums[k] || nums[i] != nums[k-1] {
		if nums[i] != nums[k-1] {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}
