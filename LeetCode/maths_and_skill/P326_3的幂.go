package main

// 数学
// 力扣：https://leetcode.cn/problems/power-of-three/

// 如果 n 是 3 的幂，即 n=3^k，那么 3^k 必然是 3^19=1162261467 的因子。
// 时间复杂度：O(1)。
// 空间复杂度：O(1)。
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
