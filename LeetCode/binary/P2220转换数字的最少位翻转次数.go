package main

import "math/bits"

// 位运算
// 力扣：https://leetcode.cn/problems/minimum-bit-flips-to-convert-number/description/

func minBitFlips(start int, goal int) int {
	// 先异或，然后找1的个数
	return bits.OnesCount(uint(start ^ goal))
}
