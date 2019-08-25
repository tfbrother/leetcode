package _8

import "sort"

// https://leetcode.com/problems/4sum/

func fourSum_err(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret, i, j, l, r, sum := make([][]int, 0, 3), 0, 0, 0, 0, 0

	for ; i < len(nums)-3; i++ {
		// TODO 此处这样优化为何逻辑有错？还没思考清楚。为何[15][16]这样优化就没有错呢？
		if nums[i] > target && i > 3 {
			break
		}

		for i > 0 && i < len(nums)-3 && nums[i] == nums[i-1] {
			i++
		}

		for j = i + 1; j < len(nums)-2; j++ {
			// 此处这样优化为何逻辑有错？还没思考清楚。借鉴的[15]优化思路。
			if nums[j] > target-nums[i] && j > i+3 {
				break
			}

			for j > i+1 && j < len(nums)-2 && nums[j] == nums[j-1] {
				j++
			}

			l, r = j+1, len(nums)-1
			for l < r {
				sum = nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					r--
					l++

					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}

		}
	}

	return ret
}

/*
Runtime: 8 ms, faster than 86.62% of Go online submissions for 4Sum.
Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for 4Sum.
*/
// 参考[15]，O(N^3)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret, i, j, l, r, sum := make([][]int, 0, 3), 0, 0, 0, 0, 0

	for ; i < len(nums)-3; i++ {
		for i > 0 && i < len(nums)-3 && nums[i] == nums[i-1] {
			i++
		}

		for j = i + 1; j < len(nums)-2; j++ {
			for j > i+1 && j < len(nums)-2 && nums[j] == nums[j-1] {
				j++
			}

			l, r = j+1, len(nums)-1
			for l < r {
				sum = nums[i] + nums[j] + nums[l] + nums[r]
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					r--
					l++

					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}

		}
	}

	return ret
}

func fourSum_en(nums []int, target int) [][]int {
	o, l := [][]int{}, len(nums)
	sort.Ints(nums)
	for i := 0; i < l-3; i++ {
		ni := nums[i]
		if i != 0 && ni == nums[i-1] {
			continue
		}
		if ni+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if ni+nums[l-1]+nums[l-2]+nums[l-2] < target {
			continue
		}
		for j := i + 1; j < l-2; j++ {
			nj := nums[j]
			if j != i+1 && nj == nums[j-1] {
				continue
			}
			if ni+nj+nums[j+1]+nums[j+2] > target {
				break
			}
			if ni+nj+nums[l-1]+nums[l-2] < target {
				continue
			}
			b, e := j+1, l-1
			for b < e {
				t := ni + nj + nums[b] + nums[e]
				if t == target {
					o, b = append(o, []int{ni, nj, nums[b], nums[e]}), b+1
					for b < e && nums[b] == nums[b-1] {
						b++
					}
				} else if t > target {
					e--
				} else {
					b++
				}
			}
		}
	}
	return o
}
