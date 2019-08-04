package _48

// =========================================================================================
// ========================================第一题============================================
// =========================================================================================
// https://leetcode-cn.com/contest/weekly-contest-148/problems/decrease-elements-to-make-array-zigzag/
func movesToMakeZigzag(nums []int) int {
	num1, num2, prev := 0, 0, 0
	// 先使得每个奇数索引对应的元素都大于相邻的元素
	for i := 1; i < len(nums); i += 2 { // i为奇数
		if nums[i-1]-prev >= nums[i] {
			num1 += nums[i-1] - nums[i] + 1 - prev
		}

		// 此处nums[i+1]会被重复的减少
		if i+1 < len(nums) && nums[i+1] >= nums[i] {
			prev = nums[i+1] - nums[i] + 1
			num1 += prev
		} else {
			prev = 0
		}
	}

	// 先使得每个偶数索引对应的元素都大于相邻的元素
	prev = 0
	for i := 0; i < len(nums); i += 2 { // i为偶数
		if i-1 > 0 && nums[i-1]-prev >= nums[i] {
			num2 += nums[i-1] - nums[i] + 1 - prev
		}

		if i+1 < len(nums) && nums[i+1] >= nums[i] {
			prev = nums[i+1] - nums[i] + 1
			num2 += prev
		} else {
			prev = 0
		}
	}

	if num1 > num2 {
		num1 = num2
	}

	return num1
}

// =========================================================================================
// ========================================第二题============================================
// =========================================================================================
// see https://leetcode-cn.com/contest/weekly-contest-148/problems/binary-tree-coloring-game/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题目考察的究竟是什么？还没搞清楚。
// 本质就是求x的相邻结点（父亲，兄弟结点），然后比较两个数的结点个数
// 分别求：x结点左右子树的结点树木，以及包含x结点路径上所有结点的数目。比较这三颗树的数目即可。
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	if root.Val == x {
		leftCount, rightCount := countNode(root.Left), countNode(root.Right)
		if leftCount != rightCount {
			return true
		}

		return false
	}

	// 计算不包含x路径的结点数量
	count1 := countNode1(root, x)
	count2, count3 := countNode(node.Left), countNode(node.Right)

	if count1 > count2+count3+1 || count2 > count1+count3+1 || count3 > count1+count2+1 {
		return true
	}

	return false
}

var node *TreeNode

func countNode(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNode(root.Left) + countNode(root.Right) + 1
}

// 计算不包含x路径的结点数量
func countNode1(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}

	if root.Val == x {
		node = root
		return 0
	}

	return countNode1(root.Left, x) + countNode1(root.Right, x) + 1
}

// =========================================================================================
// ========================================第三题============================================
// =========================================================================================
// see https://leetcode-cn.com/contest/weekly-contest-148/problems/snapshot-array/
type SnapshotArray struct {
	snapId, count int
	data          []map[int]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{count: length, data: make([]map[int]int, length, length)}
}

func (this *SnapshotArray) Set(index int, val int) {
	if this.data[index] == nil {
		this.data[index] = make(map[int]int)
	}
	this.data[index][this.snapId] = val
}

func (this *SnapshotArray) Snap() int {
	ret := this.snapId
	this.snapId++
	return ret
}

// TODO 此处是否有优化空间？
func (this *SnapshotArray) Get(index int, snap_id int) int {
	if snap_id >= this.snapId {
		snap_id = this.snapId
	}

	for i := snap_id; i >= 0; i-- {
		if _, ok := this.data[index][i]; ok {
			return this.data[index][i]
		}
	}
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

// =========================================================================================
// ========================================第四题============================================
// =========================================================================================
// see https://leetcode-cn.com/contest/weekly-contest-148/problems/longest-chunked-palindrome-decomposition/
// 应该是考察的回溯算法，然后求最大值。
func longestDecomposition(text string) int {
	tByte := []byte(text)

	//fmt.Println(tByte)
	return longestDecom(tByte, 0, len(text)-1)
}

func longestDecom(tByte []byte, start int, end int) int {
	var (
		freq [26][]int
		i, j int
	)
	for i = start; i <= end; i++ {
		j = int(tByte[i] - 'a')
		freq[j] = append(freq[j], i)
	}

	// 以首字母作为突破口
	j = int(tByte[start] - 'a')
	count := len(freq[j]) // 首字母只有一个

	if count == 1 {
		return 1
	} else { // 则检查是否可以作为一个段
		// 寻找包含首字母的段
		for i = count - 1; i > 0; i-- {
			last := freq[j][i]
			len1 := end - last + 1

			if isPalindrome(tByte[start:start+len1], tByte[last:]) {
				if start+len1 > last-1 {
					return 2
				}

				return longestDecom(tByte, start+len1, last-1) + 2
			}
		}

		return 1
	}

}

// 检查是否是回文字符串
func isPalindrome(b1, b2 []byte) bool {
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}

	return true
}
