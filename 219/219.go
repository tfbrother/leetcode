package _19

// 解法一：查找表+滑动窗口
// TODO 该版本的错误在于对go map的底层不够熟悉
// 用例 [1,2,3,1,2,3] 2 输出true，期望false
func containsNearbyDuplicate1_err0(nums []int, k int) bool {
	numIndex, i, j := make(map[int]int, k), 0, 0

	for j <= i+k && j < len(nums) {
		// 判断nums[j]是否在前面出现过
		// 如果nums[i]之前没有出现过，那么numIndex[nums[j]]=0, numIndex[nums[j]] >= 0就为true。
		if j != i && numIndex[nums[j]] >= i {
			return true
		}

		numIndex[nums[j]] = j
		if j == i+k {
			numIndex[nums[i]] = -1
			i++ // 移动窗口
		}

		j++
	}

	return false
}

// 解法一：查找表+滑动窗口
// TODO 细节没想清楚啊
// 用例 [1,0,1,1] 1  输出false，期望true
func containsNearbyDuplicate1_err1(nums []int, k int) bool {
	numIndex, i, j := make(map[int]int, k), 0, 0

	for j < len(nums) {
		// 判断nums[j]是否在前面出现过
		if _, ok := numIndex[nums[j]]; !ok {
			numIndex[nums[j]] = j
		} else if numIndex[nums[j]] >= i {
			return true
		}

		if j >= k {
			i++ // 右移动窗口
		}

		j++
	}

	return false
}

func containsNearbyDuplicate1(nums []int, k int) bool {
	numIndex, i, j := make(map[int]int, k), 0, 0

	for j < len(nums) {
		// 判断nums[j]是否在前面出现过
		if _, ok := numIndex[nums[j]]; ok && numIndex[nums[j]] >= i {
			return true
		}

		numIndex[nums[j]] = j
		if j >= k {
			i++ // 右移动窗口
		}

		j++
	}

	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	numIndex, r := make(map[int]int, k), 0

	for r < len(nums) {
		// 判断nums[j]是否在前面出现过，以及出现的位置
		if pos, ok := numIndex[nums[r]]; ok && r-pos <= k {
			return true
		}

		numIndex[nums[r]] = r
		r++
	}

	return false
}
