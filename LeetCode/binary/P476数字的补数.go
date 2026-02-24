package main

import (
	"math/bits"
)

// 位运算
// 力扣：https://leetcode.cn/problems/number-complement/

func findComplement(num int) int {
	if num == 0 {
		return 1
	}

	// 找到和 num 的二进制位数相同的全为1的二进制，然后异或运算
	return (1<<bits.Len(uint(num)) - 1) ^ num
}
