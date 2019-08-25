package _54

/*
Runtime: 44 ms, faster than 96.70% of Go online submissions for 4Sum II.
Memory Usage: 16.6 MB, less than 100.00% of Go online submissions for 4Sum II.
*/
// 和[15][18]比的区别：
// 不存在重复的问题，因为是求的索引组合。这种问题一般用查找表解决，不需要排序
// O(N^2)
func fourSumCount(A []int, B []int, C []int, D []int) int {
	map1, count := make(map[int]int, len(A)*len(A)), 0

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			map1[A[i]+B[j]]++
		}
	}

	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			count += map1[-C[i]-D[j]]
		}
	}

	return count
}
