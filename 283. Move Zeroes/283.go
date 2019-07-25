package _83__Move_Zeroes

/*
Runtime: 64 ms, faster than 76.98% of Go online submissions for Move Zeroes.
Memory Usage: 7.7 MB, less than 87.01% of Go online submissions for Move Zeroes.
*/
// 解法一：双指针，快慢指针
func moveZeroes(nums []int) {
	// k: [0, k) 存放非零元素
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 要考虑到nums全部为非零元素的情况，就会导致自己对自己赋值。
			if i != k {
				nums[k] = nums[i]
			}
			k++
		}
	}

	for ; k < len(nums); k++ {
		nums[k] = 0
	}
}

/*
Runtime: 64 ms, faster than 76.98% of Go online submissions for Move Zeroes.
Memory Usage: 7.7 MB, less than 87.01% of Go online submissions for Move Zeroes.
*/
// 解法二：解法一基础上优化
func moveZeroes2(nums []int) {
	// k: [0, k) 存放非零元素，[k, i]存放零元素
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 要考虑到nums全部为非零元素的情况，就会导致自己对自己赋值。
			if i != k {
				nums[k] = nums[i]
				nums[i] = 0
			}
			k++
		}
	}
}
