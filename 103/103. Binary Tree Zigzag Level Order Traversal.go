package _03

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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
Memory Usage: 3 MB, less than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
*/
// 解法一：解法和102完全一样，只是在返回的时将ret的奇数项数组倒序返回。
func zigzagLevelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := make([]item, 1)
	ret := make([][]int, 1)
	//queue = append(queue, item{0, root})
	queue[0] = item{0, root}

	for len(queue) > 0 {
		q := queue[0]
		node := q.node
		if len(ret)-1 < q.level {
			ret = append(ret, []int{node.Val})
		} else {
			ret[q.level] = append(ret[q.level], node.Val)
		}

		if node.Left != nil {
			queue = append(queue, item{q.level + 1, node.Left})
		}
		if node.Right != nil {
			queue = append(queue, item{q.level + 1, node.Right})
		}

		queue = queue[1:]
	}

	for i := 1; i < len(ret); i += 2 {
		j, k := 0, len(ret[i])-1
		for j < k {
			ret[i][j], ret[i][k] = ret[i][k], ret[i][j]
			j++
			k--
		}
	}

	return ret
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
Memory Usage: 3 MB, less than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
*/
// 解法二：一次性遍历每层，这是对层序遍历的深入理解，如何获取到每层的宽度
func zigzagLevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	list, ret := make([]*item, 1), make([][]int, 0)
	list[0] = &item{level: 0, node: root}

	var top *item
	for len(list) > 0 {
		width := len(list) // 记录每层的宽度
		temp := make([]int, width)
		for i := 0; i < width; i++ {
			top = list[i]
			if top.level%2 == 0 {
				temp[i] = top.node.Val
			} else {
				temp[width-i-1] = top.node.Val
			}

			if top.node.Left != nil {
				list = append(list, &item{level: top.level + 1, node: top.node.Left})
			}
			if top.node.Right != nil {
				list = append(list, &item{level: top.level + 1, node: top.node.Right})
			}
		}

		list = list[width:]
		ret = append(ret, temp)
	}

	return ret
}
