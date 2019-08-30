package _49

import (
	"fmt"
	"math"
	"strconv"
)

// 解法一：
// 可否利用机器学习的方法，一元线性回归来解决。
// 本质就是求一元一次方程组：ax+b=y的系数解，使得拟合更多的输入数据。但是存在斜率无穷大的情况啊。
// 暴力算法：递归任意两个点组成的方程组，然后求出拟合的数量。最后获得最大的值。
// 时间复杂度：O(N^3)
func maxPoints1_err(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	var a, b, count, c int

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i != j {
				// 统计拟合的数量
				c = 0
				// 求解a，b
				if points[i][0] == points[j][0] { // 垂直于X轴
					for k := 0; k < len(points); k++ {
						if points[k][0] == points[i][0] {
							c++
						}
					}
				} else if points[i][1] == points[j][1] { // 垂直于Y轴
					for k := 0; k < len(points); k++ {
						if points[k][1] == points[i][1] {
							c++
						}
					}

				} else { // 一元一次方程组
					a = (points[i][1] - points[j][1]) / (points[i][0] - points[j][0])
					b = points[i][1] - a*points[i][0]

					for k := 0; k < len(points); k++ {
						if a*points[k][0]+b == points[k][1] {
							c++
						}
					}
				}

				if count < c {
					count = c
				}
			}
		}
	}

	return count
}

// 判断i，j，k三点是否成线。
// 先固定i，再固定j，然后循环k
// 用例：[[1,1],[1,1],[2,3]] 输出 2，期望 3
func maxPoints1_1_err(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	var count, c int

	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points); j++ {
			// 统计拟合的数量
			c = 2 // 初始直线上只有i，j两点

			if points[i][0] == points[j][0] { // 垂直于X轴
				for k := j + 1; k < len(points); k++ {
					if points[k][0] == points[i][0] {
						c++
					}
				}
			} else if points[i][1] == points[j][1] { // 垂直于Y轴
				for k := j + 1; k < len(points); k++ {
					if points[k][1] == points[i][1] {
						c++
					}
				}
			} else { // 判断三点是否成一线
				for k := j + 1; k < len(points); k++ {
					// 将除法转成乘法
					if (points[j][1]-points[i][1])*(points[j][0]-points[k][0]) == (points[j][1]-points[i][1])*(points[j][0]-points[i][0]) {
						c++
					}
				}
			}

			if count < c {
				count = c
			}
		}
	}

	return count
}

// [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]] 输出 11 期望 4
func maxPoints1_2_err(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	var count, c int

	for i := 0; i < len(points)-2; i++ {
		c = 1 // 初始直线上只有i点
		for j := i + 1; j < len(points); j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] { // i,j 重合了
				c++
				continue
			} else {
				c++ // 此时直线上增加一个点，即j点
				for k := j + 1; k < len(points); k++ {
					// 判断三点是否成一线，将除法转成乘法
					if (points[j][1]-points[i][1])*(points[j][0]-points[k][0]) == (points[j][1]-points[i][1])*(points[j][0]-points[i][0]) {
						c++
					}
				}
			}

			if count < c {
				count = c
			}
		}
	}

	return count
}

// 不使用same变量
// 用例：[[0,9],[138,429],[115,359],[115,359],[-30,-102],[230,709],[-150,-686],[-135,-613],[-60,-248],[-161,-481],[207,639],[23,79],[-230,-691],[-115,-341],[92,289],[60,336],[-105,-467],[135,701],[-90,-394],[-184,-551],[150,774]]
// 输出 19， 期望 12
func maxPoints1_3_err(points [][]int) int {
	length := len(points)
	if length <= 2 {
		return length
	}

	count, max := 0, 0
	for i := 0; i < length-2; i++ {
		for j := i + 1; j < length; j++ {
			max = 2                           // 把i,j点加入到直线中
			for k := j + 1; k < length; k++ { // 寻找剩下的点中，那些在直线[i,j]上
				a := (points[i][1] - points[j][1]) * (points[i][0] - points[k][0])
				b := (points[i][1] - points[k][1]) * (points[i][0] - points[j][0])
				if a == b {
					max++
				}
			}

			if count < max {
				count = max
			}
		}
	}

	return count
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Max Points on a Line.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Max Points on a Line.
*/
// TODO 为何需要设计same变量？
// 因为算法的核心先选出[i, j]两点构成直线，然后在剩余的点中统计落在[i, j]直线上的数量。当i,j两点相同时
// 实际不形不成直线的，就会导致剩下的点都与[i, j]形成直线，但这些都是不同的直线。
func maxPoints1(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	var count, same, different int

	for i := 0; i < len(points)-2; i++ {
		same = 1 // 记录直线上与i点相同点的个数，包含i点本身。初始就只有i点一个。
		for j := i + 1; j < len(points); j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] { // i,j 重合了
				same++
				continue
			} else {
				different = 1                          // 把j点加入直线
				for k := j + 1; k < len(points); k++ { // 寻找剩下的点中，那些在直线[i,j]上
					// 判断三点是否成一线，将除法转成乘法
					if (points[j][1]-points[i][1])*(points[j][0]-points[k][0]) == (points[j][1]-points[k][1])*(points[j][0]-points[i][0]) {
						different++
					}
				}

				if count < different+same {
					count = different + same
				}
			}
		}
		// 因为有可能所有的点都相同，此时不会执行第三个for循环，就无法对count赋值
		if count < same {
			count = same
		}
	}

	return count
}

// 解法二：
// 暴力算法的基础上优化，解决精度问题
// O(N^2)
// 查找表，键为斜率
// 解决斜率精度问题：用最大公约数，3/6==2/4==1/2
// 整数问题，都要考虑0，正负数
// 数组问题一般要考虑数组长度为0，1这两种情况。
func maxPoints2_err(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	countMap, count := make(map[string]int, len(points)>>1), 0
	for i := 0; i < len(points)-1; i++ {
		// repeat表示重复的点
		repeat, max := 0, 0
		for j := i + 1; j < len(points); j++ {
			a, b := points[i][1]-points[j][1], points[i][0]-points[j][0]
			if a == 0 && b == 0 {
				repeat++
				continue
			}

			c := gcd(a, b)
			if c != 0 {
				a /= c
				b /= c
			}

			key := strconv.Itoa(a) + "/" + strconv.Itoa(b)
			countMap[key]++
			if max < countMap[key] {
				max = countMap[key]
			}
			fmt.Println(max)
		}
		fmt.Println(repeat, max)

		if count < repeat+max+1 {
			count = repeat + max + 1
		}
	}

	return count
}

/*
Runtime: 8 ms, faster than 73.75% of Go online submissions for Max Points on a Line.
Memory Usage: 4.5 MB, less than 100.00% of Go online submissions for Max Points on a Line.
*/
// 解法二：
// 暴力算法的基础上优化，解决精度问题
// O(N^2)
// 查找表，键为斜率
// 解决斜率精度问题：用最大公约数，3/6==2/4==1/2
// 整数问题，都要考虑0，正负数
// 数组问题一般要考虑数组长度为0，1这两种情况。
func maxPoints2(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	count := 0
	for i := 0; i < len(points)-1; i++ {
		// same表示重复的点，初始直线上只有i点，只有一个点与i相同，即本身。
		countMap, same, max := make(map[string]int, len(points)>>1), 1, 0
		// TODO 此时为何j从i+1开始？因为本质就是组合问题啊。
		// 选择一条直线，然后统计经过这条直线的点的个数，取最大值。组合问题顺序没有关系的。
		for j := i + 1; j < len(points); j++ {
			a, b := points[i][1]-points[j][1], points[i][0]-points[j][0]
			if a == 0 && b == 0 {
				same++
				continue
			}

			c := gcd(a, b)
			if c != 0 {
				a /= c
				b /= c
			}

			key := strconv.Itoa(a) + "/" + strconv.Itoa(b)
			countMap[key]++
			if max < countMap[key] {
				max = countMap[key]
			}
		}

		if count < same+max {
			count = same + max
		}
	}

	return count
}

// 求a，b的最大公约数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// 本质还是求斜率，只是将除法转变成乘法了。
// 而且是用三点一线的思路
// 与自己第一版的暴力算法有思想相同，只是自己没有处理好精度问题
func maxPoints_cn(points [][]int) int {
	length := len(points)
	if length <= 2 {
		return length
	}
	max := 0
	for i := 0; i < length-2; i++ {
		same := 1
		for j := i + 1; j < length; j++ {
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				same++
				continue
			}
			cnt := same + 1
			for k := j + 1; k < length; k++ {
				if (points[j][1]-points[i][1])*(points[k][0]-points[i][0]) == (points[k][1]-points[i][1])*(points[j][0]-points[i][0]) {
					cnt++
				}
			}
			max = int(math.Max(float64(max), math.Max(float64(same), float64(cnt))))
		}
		// 因为有可能所有的点都相同，此时不会执行第三个for循环，就无法对max赋值
		if max <= 0 {
			max = same
		}
	}
	return max
}
