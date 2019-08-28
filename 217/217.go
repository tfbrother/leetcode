package _17

func containsDuplicate(nums []int) bool {
	freq := make(map[int]bool, len(nums)>>1)

	for i := 0; i < len(nums); i++ {
		if freq[nums[i]] {
			return true
		}

		freq[nums[i]] = true
	}

	return false
}
