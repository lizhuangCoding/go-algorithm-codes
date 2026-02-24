package main

// 动态规划
// 力扣：https://leetcode.cn/problems/three-steps-problem-lcci/description/

func waysToStep(n int) int {
	if n == 1 || n == 2 {
		return n
	} else if n == 3 {
		return 4
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4

	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-3] + dp[i-2] + dp[i-1]) % 1000000007
	}
	return dp[n]
}

// 空间优化：
// func waysToStep(n int) int {
// 	a := 1
// 	b := 2
// 	c := a + b + 1
// 	if n == 1 { return a}
// 	if n == 2 { return b}
// 	if n == 3 { return c}
//
// 	dp1, dp2, dp3 := 1, 2, 4
// 	for i := 4; i <= n; i++ {
// 		dp1, dp2, dp3 = dp2, dp3, (dp1+dp2+dp3)%1000000007
// 	}
// 	return dp3
// }
