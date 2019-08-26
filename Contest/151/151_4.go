package main

// 4. 餐盘栈

// 核心是思考清楚，栈满和栈空时怎么表示？如何处理？
type DinnerPlates struct {
	data                  []int // 保存每个栈的数据
	stackTops             []int // 记录每个栈的栈顶元素的下一个索引，初始为0，表示没有元素
	capacity              int   // 每个栈的容量
	writeStack, readStack int   // 记录当前写和读的位置
}

func Constructor(capacity int) DinnerPlates {
	d := DinnerPlates{capacity: capacity, data: make([]int, 200000), stackTops: make([]int, 1, 100000)}
	return d
}

func (this *DinnerPlates) Push(val int) {
	if this.stackTops[this.writeStack] == this.capacity { // 当前栈已经写满了，往右找出第一个没有满的栈
		oldWriteStack, newWriteStack := this.writeStack, len(this.stackTops)

		// 从左往右，找出第一个没有满的栈
		for i := oldWriteStack; i < len(this.stackTops); i++ {
			if this.stackTops[i] != this.capacity {
				newWriteStack = i
				break
			}
		}

		// 所有栈都满了
		if newWriteStack == len(this.stackTops) {
			this.stackTops = append(this.stackTops, 0) // 新创建一个栈，默认为空。
		}

		this.writeStack = newWriteStack
		if this.writeStack > this.readStack {
			this.readStack = newWriteStack
		}
	}

	// 写入的索引
	writeIndex := this.writeStack*this.capacity + this.stackTops[this.writeStack]
	this.data[writeIndex] = val
	this.stackTops[this.writeStack]++
}

// 从右往左第一个 非空栈顶部的值
func (this *DinnerPlates) Pop() int {
	if this.stackTops[this.readStack] == 0 { // 当前栈已经空了，往左找出第一个非空的栈
		oldReadStack, newReadStack := this.readStack, -1
		// 往左找出第一个非空的栈
		for i := oldReadStack; i >= 0; i-- {
			if this.stackTops[i] > 0 {
				newReadStack = i
				break
			}
		}

		// 所有栈都为空
		if newReadStack == -1 {
			return -1
		} else {
			this.readStack = newReadStack
			if this.writeStack > this.readStack {
				this.writeStack = newReadStack
			}
		}
	}

	this.stackTops[this.readStack]--
	readIndex := this.readStack*this.capacity + this.stackTops[this.readStack]
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
	if index < this.writeStack {
		this.writeStack = index
	}
	if index > this.readStack {
		this.readStack = index
	}

	return val
}

/*
超时解法
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
*/
