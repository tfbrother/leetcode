package _1__Container_With_Most_Water

// https://leetcode-cn.com/problems/container-with-most-water/

/*
Runtime: 12 ms, faster than 94.01% of Go online submissions for Container With Most Water.
Memory Usage: 5.6 MB, less than 89.26% of Go online submissions for Container With Most Water.
*/
// TODO 要仔细思考是如何转变成对撞指针类问题的？即如何证明对撞指针解法的正确性。
// 两线段之间形成的区域总是会受到其中较短那条长度的限制。此外，两线段距离越远，得到的面积就越大。
func maxArea(height []int) int {
	l, r, max, area := 0, len(height)-1, 0, 0

	for l < r {
		if height[l] < height[r] {
			area = (r - l) * height[l]
			l++
		} else {
			area = (r - l) * height[r]
			r--
		}

		if max < area {
			max = area
		}
	}

	return max
}
