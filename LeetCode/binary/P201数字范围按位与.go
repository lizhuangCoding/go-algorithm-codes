package main

// 按位与：&
// 力扣：https://leetcode.cn/problems/bitwise-and-of-numbers-range/description/

// 暴力：时间超限

// 二进制：
// 二进制位只存在0,1两种情况，所以如果left<right，区间中必然存在left+1,那么最低位&一下一定等于0了，然后不停的右移，一直移到两个相等为止
// 时间复杂度：O(logn)。算法的时间复杂度取决于 m 和 n 的二进制位数，由于 m≤n，因此时间复杂度取决于 n 的二进制位数。
// 空间复杂度：O(1)
func rangeBitwiseAnd(left int, right int) int {
	n := 0
	for left < right {
		// 每次向右移动一位，直到left和right相等
		left >>= 1
		right >>= 1
		n++
	}

	// 此时 left == right，我们找到了left和right的二进制从左往右有连续的n位相等，然后恢复left后面的0

	// 将相等的结果左移 n 次
	return left << n
}
