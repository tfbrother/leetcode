package _5

import "sort"

// 解法一：在2sum的基础上，用暴力算法，O(N^2)。核心是如何去重，尚未实现。
// 未AC
func threeSum1(nums []int) [][]int {
	ret, numMap, target := make([][]int, 0), make(map[int]bool), 0

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			target = 0 - nums[i] - nums[j]
			if _, ok := numMap[target]; ok {
				ret = append(ret, []int{nums[i], target, nums[j]})
			} else {
				numMap[nums[j]] = true
			}
		}
	}

	return ret
}

// 解法二：排序+对撞指针，参考[167]
// 算法思想：https://leetcode-cn.com/problems/3sum/solution/3sumpai-xu-shuang-zhi-zhen-yi-dong-by-jyd/
// TODO 核心是思考清楚怎么实现去重的？
// 去重，一般首先需要排序，拍序好了才能方便去重。其次决定使用那个，要么使用第一个，要么使用最后一个。
// 又有一点组合的思想。去重，对于相同位置的，每个重复的元素只能被使用一次。
// 对于i，l位置而言，使用的是排序好的第一个出现位置，所以使用的是和nums[i-1]、nums[l-1]比较来去重的。
// 对于r而言，使用的是排序好的最后一个出现位置，所以使用的是和nums[r+1]比较。仔细体会。
func threeSum2_err(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	ret, l, r, sum := make([][]int, 0), 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { // 优化
			break
		}

		// 用例[0,0,0,0]，会出现panic: index out of range，需要保证i < len(nums)-2才行。
		// 正确：for i > 0 && i < len(nums)-2 && nums[i] == nums[i-1] {}
		//for i > 0 && nums[i] == nums[i-1] { // 跳过重复的元素
		//	i++
		//}
		if i > 0 && nums[i] == nums[i-1] { // 跳过重复的元素
			continue
		}

		l, r = i+1, len(nums)-1

		// 对撞指针
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if sum > 0 { // 移动右指针
				r--                                 // 一定要先r--
				for l < r && nums[r] == nums[r-1] { // 跳过重复的元素，去重。
					r--
				}
			} else if sum < 0 {
				l++
				for l < r && nums[l] == nums[l-1] { // 跳过重复的元素，去重。
					l++
				}
			} else {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				// 为何此处还需要这样执行呢？
				for l < r && nums[l] == nums[l-1] { // 跳过重复的元素，去重。
					l++
				}
				for l < r && nums[r] == nums[r-1] { // 跳过重复的元素，去重。
					r--
				}
			}
		}
	}

	return ret
}

func ThreeSum2(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	ret, l, r, sum := make([][]int, 0), 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { // 优化
			break
		}

		// 此处为何不能使用for，思考
		if i > 0 && nums[i] == nums[i-1] { // 跳过重复的元素
			continue
		}

		l, r = i+1, len(nums)-1

		// 对撞指针
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if sum > 0 { // 移动右指针
				r--                                 // 一定要先r--
				for l < r && nums[r] == nums[r+1] { // 跳过重复的元素，去重。
					r--
				}
			} else if sum < 0 {
				l++
				for l < r && nums[l] == nums[l-1] { // 跳过重复的元素，去重。
					l++
				}
			} else {
				temp := [3]int{nums[i], nums[l], nums[r]}
				ret = append(ret, temp[0:])
				l++
				r--
				// 为何此处还需要这样执行呢？
				for l < r && nums[l] == nums[l-1] { // 跳过重复的元素，去重。
					l++
				}
				for l < r && nums[r] == nums[r+1] { // 跳过重复的元素，去重。
					r--
				}
			}
		}
	}

	return ret
}

func ThreeSum2_1(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	ret, l, r, sum := make([][]int, 0), 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { // 优化
			break
		}

		// 此处为何不能使用for，思考
		if i > 0 && nums[i] == nums[i-1] { // 跳过重复的元素
			continue
		}

		l, r = i+1, len(nums)-1

		// 对撞指针
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if sum > 0 { // 移动右指针
				r--                                 // 一定要先r--
				for l < r && nums[r] == nums[r+1] { // 跳过重复的元素，去重。
					r--
				}
			} else if sum < 0 {
				l++
				for l < r && nums[l] == nums[l-1] { // 跳过重复的元素，去重。
					l++
				}
			} else {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				// 为何此处还需要这样执行呢？
				for l < r && nums[l] == nums[l-1] { // 跳过重复的元素，去重。
					l++
				}
				for l < r && nums[r] == nums[r+1] { // 跳过重复的元素，去重。
					r--
				}
			}
		}
	}

	return ret
}

func ThreeSum2_2(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	ret, l, r, sum := make([][]int, 0), 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { // 优化
			break
		}

		// 此处为何不能使用for，思考
		if i > 0 && nums[i] == nums[i-1] { // 跳过重复的元素
			continue
		}

		l, r = i+1, len(nums)-1

		// 对撞指针
		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if sum > 0 { // 移动右指针
				r--
			} else if sum < 0 {
				l++
			} else {
				// temp := []int{nums[i], nums[l], nums[r]}
				// fmt.Println(cap(temp))	// output 3
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				// 为何此处还需要这样执行呢？
				for l < r && nums[l] == nums[l-1] { // 跳过重复的元素，去重。
					l++
				}
				for l < r && nums[r] == nums[r+1] { // 跳过重复的元素，去重。
					r--
				}
			}
		}
	}

	return ret
}

func ThreeSum_en(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			left := i + 1
			right := len(nums) - 1
			for left < right {
				if nums[left]+nums[right] > -nums[i] {
					right--
				} else if nums[left]+nums[right] < -nums[i] {
					left++
				} else {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					right--
					left++
				}
			}
		}
	}
	return result
}
