package main

/*
Runtime: 4 ms, faster than 96.50% of Go online submissions for Two Sum.
Memory Usage: 3.7 MB, less than 15.02% of Go online submissions for Two Sum.
*/
// 实现一，采用临时变量。
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := target - nums[i]

		if _, ok := numMap[num]; ok {
			return []int{numMap[num], i}
		}

		numMap[nums[i]] = i
	}

	return nil
}

/*
Runtime: 0 ms, faster than 100% of Go online submissions for Two Sum.
Memory Usage: 3.7 MB, less than 15.02% of Go online submissions for Two Sum.
*/
// 实现二，采用局部变量。优化内存的作用，优化cpu耗时。
func twoSum2(nums []int, target int) []int {
	var (
		numMap map[int]int
		num, j int
		ok     bool
	)

	numMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num = target - nums[i]
		if j, ok = numMap[num]; ok {
			return []int{j, i}
		}

		numMap[nums[i]] = i
	}

	return nil
}

/*
func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)

    for i :=0; i < len(nums); i++ {
        num := target - nums[i]
        fmt.Println(&num)   // 打印地址
        if _, ok := numMap[num]; ok {
            return []int{numMap[num], i}
        } else {
            fmt.Println(&ok)    // 打印地址
        }

        numMap[nums[i]] = i
    }

    return nil
}

input:
[2,7,11,15]
9
debug output:
0xc0000200b8
0xc0000200d0
0xc0000200d8

可见每次变量的地址是不一样的，在go运行时内部每次都会重新申请一个临时变量并进行赋值，这肯定比只对变量进行赋值要慢一些。
由于go运行时有自己的内存管理，假设每次是事先向操作系统申请一大堆内存(比如10M)，然后自己管理。
所以两种方法提交时内存占用率一样，是因为用临时变量后内存占用没有使用超过10M，所以内存占用显示出来就只有10M。
*/
