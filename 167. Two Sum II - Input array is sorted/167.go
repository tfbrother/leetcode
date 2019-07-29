package _67__Two_Sum_II___Input_array_is_sorted

/*
Runtime: 4 ms, faster than 96.52% of Go online submissions for Two Sum II - Input array is sorted.
Memory Usage: 3 MB, less than 100.00% of Go online submissions for Two Sum II - Input array is sorted.
*/
// 解法一：hash表，参考:[1].Two Sum的解法
// 但是该解法没有利用到numbers是有序的这个性质
// 时间：O(N)，空间O(N)
func twoSum(numbers []int, target int) []int {
	numMap := make(map[int]int)

	for i := 0; i < len(numbers); i++ {
		num := target - numbers[i]

		if _, ok := numMap[num]; ok {
			return []int{numMap[num] + 1, i + 1}
		}

		numMap[numbers[i]] = i
	}

	return nil
}

/*
Runtime: 4 ms, faster than 96.52% of Go online submissions for Two Sum II - Input array is sorted.
Memory Usage: 3 MB, less than 61.54% of Go online submissions for Two Sum II - Input array is sorted.
*/
// 解法二：二分查找法，利用到了数组有序，但是时间复杂度反而提高了
// 时间：O(NlogN)，空间：O(1)
func twoSum2(numbers []int, target int) []int {
	var indexAt int
	for i := 0; i < len(numbers); i++ {
		indexAt = binarySearch(numbers, i+1, len(numbers)-1, target-numbers[i])
		if indexAt != -1 {
			return []int{i + 1, indexAt + 1}
		}
	}

	return nil
}

// 在numsbers中二分查找val，查到返回索引位置，没有查到返回-1
func binarySearch(numbers []int, l int, r int, val int) int {
	middle := 0
	for l <= r {
		middle = l + (r-l)>>1
		if numbers[middle] == val {
			return middle
		} else if numbers[middle] > val {
			r = middle - 1
		} else {
			l = middle + 1
		}
	}

	return -1
}

// 解法三：对撞指针(双索引技术)
// 时间：O(N)，空间：O(1)
func twoSum3(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if target == numbers[l]+numbers[r] {
			return []int{l + 1, r + 1}
		} else if target < numbers[l]+numbers[r] {
			r--
		} else {
			l++
		}
	}

	return nil
}

// 扩展1：如果有多个解，请返回所有的解法。只在解法三的基础上实现
// 此时要考虑要考虑数组中是否有重复的元素了。
// 用例:[1,2,3,4,5,6,7,8,9,10], 15
func twoSum3_1(numbers []int, target int) [][]int {
	l, r, ret := 0, len(numbers)-1, make([][]int, 0)
	for l < r {
		if target == numbers[l]+numbers[r] {
			ret = append(ret, []int{l + 1, r + 1})
			// 同时进行l++，r--的前提是，数组中没有重复的元素。否则进行l++或者r--
			l++
			r--
		} else if target < numbers[l]+numbers[r] {
			r--
		} else {
			l++
		}
	}

	return ret
}

// 扩展2：返回l+r值最大的那个解
func twoSum3_2(numbers []int, target int) []int {
	l, r, ret, max := 0, len(numbers)-1, make([]int, 0), 0
	for l < r {
		if target == numbers[l]+numbers[r] {
			if max < l+r {
				max = l + r
				ret = []int{l + 1, r + 1}
			}
			// 同时进行l++，r--的前提是，数组中没有重复的元素。否则进行l++或者r--
			l++
			r--
		} else if target < numbers[l]+numbers[r] {
			r--
		} else {
			l++
		}
	}

	return ret
}
