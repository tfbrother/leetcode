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

func main() {
	str := []string{"bob,627,1973,amsterdam", "alex,387,885,bangkok", "alex,355,1029,barcelona", "alex,587,402,bangkok", "chalicefy,973,830,barcelona", "alex,932,86,bangkok", "bob,188,989,amsterdam"}
	fmt.Println(invalidTransactions(str))
}
