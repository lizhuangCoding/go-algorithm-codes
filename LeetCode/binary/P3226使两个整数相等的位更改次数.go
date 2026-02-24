package main

import "math/bits"

// 位运算
// 力扣：https://leetcode.cn/problems/number-of-bit-changes-to-make-two-integers-equal/

func minChanges(n int, k int) int {
	// 只能选择 n 的 二进制表示 中任意一个值为 1 的位，并将其改为 0。
	if n&k != k {
		return -1
	}

	// 返回 n ^ k 的二进制的1的个数
	return bits.OnesCount(uint(n ^ k))
}
