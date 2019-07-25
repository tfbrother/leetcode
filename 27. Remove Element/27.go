package _7__Remove_Element

// 和283题完全一样。
func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != k {
				nums[k] = nums[i]
			}
			k++
		}
	}

	return k
}
