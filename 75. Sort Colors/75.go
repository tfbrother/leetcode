package _5__Sort_Colors

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Sort Colors.
*/
// 解法一：计数排序
func sortColors(nums []int) {
	var counts = [3]int{0, 0, 0}
	for i := 0; i < len(nums); i++ {
		counts[nums[i]]++
	}
	k := 0
	for i := 0; i < len(counts); i++ {
		for j := 0; j < counts[i]; j++ {
			nums[k] = i
			k++
		}
	}
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
Memory Usage: 2.3 MB, less than 64.15% of Go online submissions for Sort Colors.
*/
// 解法二：三指针--三路快排思想
func sortColors2(nums []int) {
	// nums[0...zero]存放0，nums[zero+1...i-1]存放1，nums[i..two-1]未遍历元素，nums[two..]存放2
	// zero和two的初值一定要搞清楚
	zero, two, i := -1, len(nums), 0

	for i < two { // 这里也要想清楚，结束条件是i < two
		if nums[i] == 0 {
			zero++
			// nums[zero], nums[i] = 0, 1
			// 注意此时并不能保证nums[zero] == 1了，因为此时可能i == zero
			nums[zero], nums[i] = 0, nums[zero]
			i++
		} else if nums[i] == 1 {
			i++
		} else { // nums[i] == 2
			two--
			nums[i], nums[two] = nums[two], 2
		}
	}
}
