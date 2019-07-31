package _25__Maximum_Size_Subarray_Sum_Equals_k

// 解法一：滑动窗口技术
// 第一版实现中，超出时间限制了。没有考虑到nums[i]>s的情况。
func minSubArrayLen1(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		if nums[0] == s {
			return 1
		}

		return 0
	}

	i, j, min, sum, length := 0, 1, 0, nums[0], 0
	for j < len(nums) {
		if sum+nums[j] == s {
			length = j - i + 1
			if min > length {
				min = length
			} else if min == 0 { // min == 0放在else中，也是一种优化。
				min = length
			}
		} else if sum+nums[j] > s {
			sum -= nums[i]
			i++
		} else {
			sum += nums[j]
			j++
		}
	}

	return min
}

// 第二版实现中，超出时间限制了。没有考虑到nums[i]>s的情况。
func minSubArrayLen2(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		if nums[0] == s {
			return 1
		}

		return 0
	}

	// i,j表示窗口的左右边界，sum表示nums[i...j-1]的和。
	i, j, min, sum, length := 0, 0, 0, 0, 0
	for j < len(nums) { // 循环终止条件
		if nums[j] == s {
			return 1
		} else if nums[j] > s { // 排除nums[j]，同时更新窗口
			j++
			i = j
			sum = 0
		} else if sum+nums[j] == s { // 此时需要移动窗口，i++，不能同时i++，j++，因为可能nums[i]==0
			length = j - i + 1
			if min > length {
				min = length
			} else if min == 0 { // min == 0放在else中，也是一种优化。
				min = length
			}
			i++
		} else if sum+nums[j] > s {
			sum -= nums[i]
			i++
		} else {
			sum += nums[j]
			j++
		}
	}

	return min
}

// 解法三，正确
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		if nums[0] == s {
			return 1
		}

		return 0
	}

	// i,j表示窗口的左右边界，sum表示nums[i...j-1]的和。
	i, j, min, sum, length := 0, 0, 0, 0, 0
	for j < len(nums) { // 循环终止条件
		if nums[j] == s {
			return 1
		} else if nums[j] > s { // 排除nums[j]，同时更新窗口
			j++
			i = j
			sum = 0
		} else if sum+nums[j] == s { // 此时需要移动窗口，i++，不能同时i++，j++，因为可能nums[i]==0
			length = j - i + 1
			if min > length {
				min = length
			} else if min == 0 { // min == 0放在else中，也是一种优化。
				min = length
			}
			sum -= nums[i]
			i++
		} else if sum+nums[j] > s {
			sum -= nums[i]
			i++
		} else { // < s
			sum += nums[j]
			j++
		}
	}

	return min
}
