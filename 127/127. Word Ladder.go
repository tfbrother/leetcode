package _27

/*
Runtime: 196 ms, faster than 39.36% of Go online submissions for Word Ladder.
Memory Usage: 5.7 MB, less than 50.00% of Go online submissions for Word Ladder.
*/
// 图论，最短路径
// 只相差一个字母的单词，就形成了一条边
// TODO 核心是如何判断两个字符串只相差一个字符
func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue, visted := make([][2]int, 1), make([]bool, len(wordList))
	queue[0] = [2]int{-1, 1}

	for len(queue) > 0 {
		index, step, word := queue[0][0], queue[0][1], ""
		if index == -1 {
			word = beginWord
		} else {
			word = wordList[index]
		}

		nearlys := getNearlyWord(word, wordList)
		//fmt.Println(nearlys, word)
		for i := 0; i < len(nearlys); i++ {
			if !visted[nearlys[i]] {
				if wordList[nearlys[i]] == endWord {
					return step + 1
				}
				queue = append(queue, [2]int{nearlys[i], step + 1})
				visted[nearlys[i]] = true
			}
		}
		queue = queue[1:]
	}

	return 0
}

// 暴力解法：获取只相差一个字符的单词索引列表
func getNearlyWord(word string, wordList []string) []int {
	ret := make([]int, 0)
	for i := 0; i < len(wordList); i++ {
		diff := 0
		for j := 0; j < len(word); j++ {
			if word[j] != wordList[i][j] {
				diff++
				if diff > 1 {
					break
				}
			}
		}

		if diff == 1 {
			ret = append(ret, i)
		}
	}

	return ret
}
