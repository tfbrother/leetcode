package _20

// 解法一：滑动窗口+暴力算法, O(N^2)
// 当k>=len(nums)时，该解法会固定返回false。
// 因为i,j是不同的索引，所以k>0
func containsNearbyAlmostDuplicate1_err(nums []int, k int, t int) bool {
	if len(nums) == 0 || t < 0 || k <= 0 {
		return false
	}

	for i := 0; i < len(nums)-k; i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if abs(nums[j]-nums[i]) <= t {
				return true
			}
		}
	}

	return false
}

// 解法一：滑动窗口+暴力算法, O(N^2)
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	if len(nums) == 0 || t < 0 || k <= 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if abs(nums[j]-nums[i]) <= t {
				return true
			}
		}
	}

	return false
}

// 解法一：滑动窗口+暴力算法, O(N^2)
// 当k等于0时会出错，因为l，r的初始值就默认了k至少为1。
// 因为i,j是不同的索引，所以k>0
// 归根结底还是没有把滑动窗口的搞清楚，该滑动窗口的大小是不固定，在[1,k]之间
func containsNearbyAlmostDuplicate2_err(nums []int, k int, t int) bool {
	if len(nums) == 0 || t < 0 || k < 0 {
		return false
	}

	l, r := 0, 1
	for r < len(nums) {
		if abs(nums[r]-nums[l]) <= t {
			return true
		}

		if r >= k { // TODO 此处逻辑错误，会使得r>=k后，后面的窗口大小恒为1了。
			l++
			r = l + 1
		} else {
			r++
		}
	}

	return false
}

// 解法一：滑动窗口+暴力算法, O(N^2)
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	if len(nums) == 0 || t < 0 || k <= 0 {
		return false
	}

	l, r := 0, 1
	for r < len(nums) {
		if abs(nums[r]-nums[l]) <= t {
			return true
		}

		if r-l >= k || r == len(nums)-1 { // 维持窗口的大小在[1,k]之间
			l++
			r = l + 1
		} else {
			r++
		}
	}

	return false
}

// 解法二：滑动窗口+桶排序，O(N)
// 当t==0时就是[219]题
// TODO 核心是思考，如何用桶排序实现绝对值差的问题，以及为何需要处理负数问题。
func containsNearbyAlmostDuplicate3(nums []int, k int, t int) bool {
	if len(nums) == 0 || t < 0 || k < 0 {
		return false
	}

	windowMap, r, prod := make(map[int]int), 0, 0
	for r < len(nums) {
		// -4/5=0，4/5=0，4与-4之间却相差8.
		prod = (nums[r] - -1>>63) / (t + 1)
		if _, ok := windowMap[prod]; ok {
			return true
		}

		if v, ok := windowMap[prod-1]; ok && nums[r]-v <= t {
			return true
		}

		if v, ok := windowMap[prod+1]; ok && v-nums[r] <= t {
			return true
		}

		windowMap[prod] = nums[r]
		if r >= k {
			delete(windowMap, (nums[r-k]- -1>>63)/(t+1))
		}
		r++
	}

	return false
}

// 桶排序+滑动窗口
// TODO 区别就是在与负数的处理。采用的方式就是将负数的商全部减1。确实是不错的技巧
func containsNearbyAlmostDuplicate_cn(nums []int, k int, t int) bool {
	//桶排序，t为桶元素个数；遍历数组，保存每个元素所在的桶的位置
	//map[idx] = nums[i]；然后看桶内，桶前，桶后这三个桶的位置是否有元素，有的话则返回true；
	//当滑动窗口超过了k，则删除i-k之前的元素（获取桶位置，然后再看该位置里面是否有元素）
	//要找出窗口k的元素，建立一个个桶；同一个桶内的元素，差值肯定是满足；当然这里要考虑k这个因素，当元素的值超过桶k，需要额外处理
	//不同的桶内，则需要考虑两个桶是否相邻，相邻的桶内也有可能存在窗口k的元素。
	if t < 0 {
		return false
	}
	bmap := make(map[int]int)
	w := t + 1
	for i := 0; i < len(nums); i++ {
		idx := getBucketId(nums[i], w)

		if _, ok := bmap[idx]; ok {
			return true
		}

		if val, ok := bmap[idx-1]; ok {
			if nums[i]-val <= t {
				return true
			}
		}

		if val, ok := bmap[idx+1]; ok {
			if val-nums[i] <= t {
				return true
			}
		}

		bmap[idx] = nums[i]

		if i >= k {
			idx2 := getBucketId(nums[i-k], w)
			delete(bmap, idx2)
		}
	}
	return false
}

func getBucketId(m int, w int) int {
	if m < 0 { //负数处理
		return m/w - 1
	}
	return m / w
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 其实整型溢出还没有考虑
