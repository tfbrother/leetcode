package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 1. 查询无效交易
// transactions.length <= 1000
// 每笔交易 transactions[i] 按 "{name},{time},{amount},{city}" 的格式进行记录
// 每个交易名称 {name} 和城市 {city} 都由小写英文字母组成，长度在 1 到 10 之间
// 每个交易时间 {time} 由一些数字组成，表示一个 0 到 1000 之间的整数
// 每笔交易金额 {amount} 由一些数字组成，表示一个 0 到 2000 之间的整数

// TODO 核心是解决：它和另一个城市中同名的另一笔交易相隔不超过 60 分钟（包含 60 分钟整）
// 卧槽：最后采用暴力算法，才解决了。
// 题意没有理解清楚：
// 用例：["alice,20,800,mtv","alice,50,1200, barcelona"]
// 期望：["alice,20,800,mtv","alice,50,1200, barcelona"]
func invalidTransactions(transactions []string) []string {
	ret := make([]string, 0, len(transactions)>>1)
	tranMap := make(map[string][][]int)
	citySet := make(map[string]int)
	var (
		name, city   string
		time, amount int
	)

	for i := 0; i < len(transactions); i++ {
		tranInfo := splitTran(transactions[i])
		name = tranInfo[0]
		time, _ = strconv.Atoi(tranInfo[1])
		amount, _ = strconv.Atoi(tranInfo[2])
		city = tranInfo[3]
		if _, ok := citySet[city]; !ok {
			citySet[city] = len(citySet)
		}

		if _, ok1 := tranMap[name]; ok1 {
			tranMap[name] = append(tranMap[name], []int{citySet[city], i, amount, time})
		} else {
			tranMap[name] = [][]int{[]int{citySet[city], i, amount, time}}
		}
	}

	retIndexes := make([]int, 0, len(transactions))
	// 处理时间不满足条件的元素
	for _, item := range tranMap {
		// 按照时间排序
		sort.Slice(item, func(i, j int) bool {
			return item[i][3] < item[j][3]
		})

		// 暴力查找
		for i := 0; i < len(item); i++ {
			if item[i][2] > 1000 { // 金额超过1000
				retIndexes = append(retIndexes, item[i][1])
			}
			// 暴力查找，时间不超过60分钟的
			for j := i - 1; j >= 0; j-- {
				if item[i][0] == item[j][0] { // 城市相同，继续
					continue
				}
				if item[i][3]-item[j][3] <= 60 {
					retIndexes = append(retIndexes, item[i][1], item[j][1])
				} else {
					break
				}
			}
		}
	}

	// 去重retIndexes
	used := make(map[int]bool, len(retIndexes))
	for i := 0; i < len(retIndexes); i++ {
		if _, ok := used[retIndexes[i]]; !ok {
			ret = append(ret, transactions[retIndexes[i]])
			used[retIndexes[i]] = true
		}
	}

	return ret
}

func splitTran(str string) []string {
	l, ret := 0, make([]string, 0, 4)
	for i := 0; i < len(str); i++ {
		if str[i] == ',' {
			ret = append(ret, string(str[l:i]))
			l = i + 1
		}
	}

	ret = append(ret, string(str[l:]))
	return ret
}

// 2. 比较字符串最小字母出现频次
func numSmallerByFrequency(queries []string, words []string) []int {
	freqWords, ret := make([]int, 0, len(words)), make([]int, 0, len(queries))
	for i := 0; i < len(words); i++ {
		freqWords = append(freqWords, getFreq([]byte(words[i])))
	}

	for j := 0; j < len(queries); j++ {
		freq, count := getFreq([]byte(queries[j])), 0
		for k := 0; k < len(freqWords); k++ {
			if freqWords[k] > freq {
				count++
			}
		}

		ret = append(ret, count)
	}

	return ret
}

// f函数
func getFreq(bs []byte) int {
	var freqs [26]int
	minC := 27

	for i := 0; i < len(bs); i++ {
		j := int(bs[i] - 'a')
		if minC > j {
			minC = j
		}
		freqs[j]++
	}

	return freqs[minC]
}

// 3. 从链表中删去总和值为零的连续节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除当前结点，或者不删除当前结点，递归处理
// 反复删去链表中由 总和 值为 0 的连续节点组成的序列
// 如何实现反复删除？
func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	head, flag := removeZero(head)
	if flag {
		head = removeZeroSumSublists(head)
	} else {
		child := removeZeroSumSublists(head.Next)
		head.Next = child
	}

	return head
}

// 从head开始删除和连续为0的，成功则返回
func removeZero(head *ListNode) (newHead *ListNode, flag bool) {
	if head == nil {
		return nil, false
	}

	node := head
	sum := 0
	for node != nil {
		sum += node.Val
		if sum == 0 {
			return node.Next, true
		}
		node = node.Next
	}

	return head, false
}

// 4. 餐盘栈

// 核心是思考清楚，栈满和栈空时怎么表示？如何处理？
// 超出时间限制了。
type DinnerPlates struct {
	data      []int // 保存每个栈的数据
	stackTops []int // 记录每个栈的栈顶元素的下一个索引，初始为0，表示没有元素
	capacity  int   // 每个栈的容量
}

func Constructor(capacity int) DinnerPlates {
	d := DinnerPlates{capacity: capacity, data: make([]int, 200000), stackTops: make([]int, 0, 100000)}
	return d
}

func (this *DinnerPlates) Push(val int) {
	// 从左往右，找出第一个没有满的栈
	writeStack := len(this.stackTops)
	for i := 0; i < len(this.stackTops); i++ {
		if this.stackTops[i] != this.capacity {
			writeStack = i
			break
		}
	}

	// 所有栈都满了
	if writeStack == len(this.stackTops) {
		this.stackTops = append(this.stackTops, 0) // 新创建一个栈，默认为空。
	}

	// 写入的索引
	writeIndex := writeStack*this.capacity + this.stackTops[writeStack]
	this.data[writeIndex] = val
	this.stackTops[writeStack]++
}

// 从右往左第一个 非空栈顶部的值
func (this *DinnerPlates) Pop() int {
	// 从右往左第一个 非空栈顶部的值
	readStack := -1
	for i := len(this.stackTops) - 1; i >= 0; i-- {
		if this.stackTops[i] > 0 {
			readStack = i
			break
		}
	}

	// 所有栈都为空
	if readStack == -1 {
		return -1
	}

	this.stackTops[readStack]--
	readIndex := readStack*this.capacity + this.stackTops[readStack]
	val := this.data[readIndex]

	return val
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if this.stackTops[index] == 0 {
		return -1
	}
	this.stackTops[index]--
	readIndex := index*this.capacity + this.stackTops[index]
	val := this.data[readIndex]

	return val
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */

func main() {
	str := []string{"bob,627,1973,amsterdam", "alex,387,885,bangkok", "alex,355,1029,barcelona", "alex,587,402,bangkok", "chalicefy,973,830,barcelona", "alex,932,86,bangkok", "bob,188,989,amsterdam"}
	fmt.Println(invalidTransactions(str))
}
