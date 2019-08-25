package _47

// n最大为500，可见可以使用暴力算法
// 还要根据坐标的范围，决定是否需要考虑整型溢出
// 该版本有两个错误：
// 	1. i != j时，应该跳过继续下一个循环，而不是退出循环
//	2. 每个i都应该单独一个distMap，比如i=0时，与它距离1的点有2个；i=2时，与它距离1的点也有2个。
// 	   count=4*(4-1)=12 还是 count=2*(2-1)+2*(2-1)=4 呢？仔细思考。
//  用例[[0,0],[1,0],[-1,0],[0,1],[0,-1]]输出26，期望20
func numberOfBoomerangs_err(points [][]int) int {
	distMap, count, dist := make(map[int]int), 0, 0

	// 记录每一个点与其它点的距离
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points) && i != j; j++ {
			dist = (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			distMap[dist]++
		}
	}

	for _, c := range distMap {
		count += c * (c - 1)
	}

	return count
}

/*
Runtime: 156 ms, faster than 69.57% of Go online submissions for Number of Boomerangs.
Memory Usage: 7.7 MB, less than 100.00% of Go online submissions for Number of Boomerangs.
*/
// n最大为500，可见可以使用暴力算法
// 还要根据坐标的范围，决定是否需要考虑整型溢出
func numberOfBoomerangs(points [][]int) int {
	var count, dist int

	// 记录每一个点与其它点的距离
	for i := 0; i < len(points); i++ {
		distMap := make(map[int]int, len(points)>>1)
		for j := 0; j < len(points); j++ {
			if i != j {
				dist = (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
				distMap[dist]++
			}
		}

		for _, c := range distMap {
			count += c * (c - 1)
		}
	}

	return count
}

// 利用数学知识，优化排列的求解。不过还没看理解清楚。
func numberOfBoomerangs_en(points [][]int) int {
	result := 0
	for i, pointI := range points {
		hashMap := make(map[int]int, len(points))
		for j, pointJ := range points {
			if i == j {
				continue
			}
			distance := (pointJ[0]-pointI[0])*(pointJ[0]-pointI[0]) + (pointJ[1]-pointI[1])*(pointJ[1]-pointI[1])

			if n, exist := hashMap[distance]; exist {
				result += n * 2
				hashMap[distance]++
			} else {
				hashMap[distance] = 1
			}
		}
	}

	return result
}
