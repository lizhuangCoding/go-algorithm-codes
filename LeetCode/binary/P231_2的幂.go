package main

// 数学、二进制
// 力扣：https://leetcode.cn/problems/power-of-two/

// 方法一：二进制
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// 方法二：判断是否为最大 2 的幂的约数
// 在题目给定的 32 位有符号整数的范围内（ n <= 2^31 - 1），最大的 2 的幂为 2^30=1073741824。我们只需要判断 n 是否是 2^30 的约数即可。
// func isPowerOfTwo(n int) bool {
// 	m := 1 << 30
// 	return n > 0 && m%n == 0
// }
