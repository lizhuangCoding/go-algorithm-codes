package main

import "math/bits"

// 位运算
// 力扣：https://leetcode.cn/problems/smallest-number-with-all-set-bits/description/

func smallestNumber(n int) int {
	// 计算 n 的二进制长度 m，返回长为 m 的全为 1 的二进制数，也就是 2^m -1
	return 1<<bits.Len(uint(n)) - 1
}
