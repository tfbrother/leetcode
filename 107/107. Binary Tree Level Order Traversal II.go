package _07

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type item struct {
	level int       // 层数
	node  *TreeNode // 结点
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
*/
// 解法一：在[102]的基础上，对结果进行倒序
func levelOrderBottom1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ret := levelOrder(root)
	// 反转数组（对撞指针）
	l, r := 0, len(ret)-1
	for l < r {
		ret[l], ret[r] = ret[r], ret[l]
		l++
		r--
	}

	return ret
}

func levelOrder(root *TreeNode) [][]int {
	list, ret := make([]*item, 1), make([][]int, 0)
	list[0] = &item{level: 0, node: root}

	var top *item
	for len(list) > 0 {
		top, list = list[0], list[1:]
		if len(ret) == top.level {
			ret = append(ret, []int{top.node.Val})
		} else {
			ret[top.level] = append(ret[top.level], top.node.Val)
		}
		if top.node.Left != nil {
			list = append(list, &item{level: top.level + 1, node: top.node.Left})
		}
		if top.node.Right != nil {
			list = append(list, &item{level: top.level + 1, node: top.node.Right})
		}
	}

	return ret
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Binary Tree Level Order Traversal II.
*/
// 解法二：提前获取到树的高度，然后存储的时候逆序存储
// 严格来说，比解法一更优
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	h := getHeight(root)
	list, ret := make([]*item, 1), make([][]int, h)
	list[0] = &item{level: 1, node: root}

	var top *item
	for len(list) > 0 {
		top, list = list[0], list[1:]
		ret[h-top.level] = append(ret[h-top.level], top.node.Val)

		if top.node.Left != nil {
			list = append(list, &item{level: top.level + 1, node: top.node.Left})
		}
		if top.node.Right != nil {
			list = append(list, &item{level: top.level + 1, node: top.node.Right})
		}
	}

	return ret
}

// 获取树的高度，根节点的高度为1
// DFS
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := getHeight(root.Left), getHeight(root.Right)
	if left >= right {
		return left + 1
	}

	return right + 1
}
