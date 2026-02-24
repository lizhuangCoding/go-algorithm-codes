package main

// 动态规划：简单
// 力扣：https://leetcode.cn/problems/n-th-tribonacci-number/description/

func tribonacci(n int) int {
	dp := make([]int, n+1)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}

	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}
	return dp[n]
}

// 空间优化：
// func tribonacci(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	if n <= 2 {
// 		return 1
// 	}
// 	p, q, r, s := 0, 0, 1, 1
// 	for i := 3; i <= n; i++ {
// 		p = q
// 		q = r
// 		r = s
// 		s = p + q + r
// 	}
// 	return s
// }
