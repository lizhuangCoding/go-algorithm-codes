package main

// 位运算：& ^
// 力扣：https://leetcode.cn/problems/single-number-iii/description/

// 方法一：map
// 省略...

// 方法二：位运算
// 题解：力扣官方：https://leetcode.cn/problems/single-number-iii/solutions/587516/zhi-chu-xian-yi-ci-de-shu-zi-iii-by-leet-4i8e/
// lsb := xorSum & -xorSum：找到最低有效位
// xorSum & -xorSum 计算 xorSum 的最低有效位，也就是 xorSum 中最右边的为 1 的位。这一位是 a 和 b 在二进制表示中不同的那一位，因为 a ^ b 的结果在某一位上为 1，意味着 a 和 b 在这一位上必然是不同的，一个是 0，一个是 1。
// lsb 是用来将两个数字 a 和 b 分成两类的依据，它是二进制中最右边的 1，表示在这一位上 a 和 b 有不同的值。
func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}
