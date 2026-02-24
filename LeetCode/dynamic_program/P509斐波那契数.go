package main

// 动态规划
// 力扣：https://leetcode.cn/problems/fibonacci-number/description/

func fib(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	if n == 0 {
		return 0
	}
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 空间优化：
// func fib(n int) int {
// 	if n < 2 {
// 		return n
// 	}
// 	a, b, c := 0, 1, 0
// 	for i := 1; i < n; i++ {
// 		c = a + b
// 		a, b = b, c
// 	}
// 	return c
// }
