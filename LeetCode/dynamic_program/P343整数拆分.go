package main

// 动态规划
// 力扣：https://leetcode.cn/problems/integer-break/

func integerBreak(n int) int {
	dp := make([]int, n+1) // dp[i]：分拆数字i，可以得到的最大乘积为dp[i]。
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			// i可以差分为i-j和j。由于需要最大值，故需要通过j遍历所有存在的值，取其中最大的值作为当前i的最大值，在求最大值的时候，一个是j与i-j相乘，一个是j与dp[i-j].
			// j * (i - j) 是单纯的把整数拆分为两个数相乘，而j * dp[i - j]是拆分成两个以及两个以上的个数相乘。
			dp[i] = maxIntegerBreak(dp[i], maxIntegerBreak(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func maxIntegerBreak(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
