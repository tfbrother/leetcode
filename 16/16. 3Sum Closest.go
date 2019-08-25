package _6

import (
	"math"
	"sort"
)

// 解题思路和[15]一摸一样，但是不需要去重了。
func threeSumClosest1_err(nums []int, target int) int {
	sort.Ints(nums)
	min, i, l, r, closest, diff, sum := math.MaxInt64, 0, 0, 0, 0, 0, 0

	for ; i < len(nums)-2; i++ {
		// 所有nums[0]都大于target这种情况，没有考虑到。
		if nums[i] > target {
			break
		}

		l, r = i+1, len(nums)-1
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			diff = sum - target
			if diff > 0 {
				r--
			} else if diff < 0 {
				l++
				diff = -diff
			} else {
				return target
			}

			if min > diff {
				min = diff
				closest = sum
			}
		}
	}

	return closest
}

/*
Runtime: 4 ms, faster than 98.38% of Go online submissions for 3Sum Closest.
Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for 3Sum Closest.
*/
// 解题思路和[15]一摸一样，但是不需要去重了。
func threeSumClosest1_1(nums []int, target int) int {
	sort.Ints(nums)
	min, i, l, r, closest, diff, sum := math.MaxInt64, 0, 0, 0, 0, 0, 0

	for ; i < len(nums)-2; i++ {
		// 所有nums[i]都大于target这种情况，没有考虑到。
		// 需要增加i > 2这种情况，表示至少处理了前三个元素了
		if nums[i] > target && i > 2 {
			break
		}

		l, r = i+1, len(nums)-1
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			diff = sum - target
			if diff > 0 {
				r--
			} else if diff < 0 {
				l++
				diff = -diff
			} else {
				return target
			}

			if min > diff {
				min = diff
				closest = sum
			}
		}
	}

	return closest
}
