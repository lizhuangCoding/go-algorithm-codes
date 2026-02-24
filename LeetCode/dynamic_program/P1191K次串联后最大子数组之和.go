package main

// 动态规划 一、入门DP 1.3 最大子数组和（最大子段和）
// 力扣：https://leetcode.cn/problems/k-concatenation-maximum-sum/description/

func kConcatenationMaxSum(arr []int, k int) int {
	mod := 1_000_000_007
	sum := 0
	for _, num := range arr {
		sum = (sum + num) % mod
	}

	if k >= 2 {
		arr = append(arr, arr...)
	}
	// maxSum记录当前最大子数组和
	n, maxSum := len(arr), 0
	dp := make([]int, n+1) // dp[i] 表示以 arr[i-1] 结尾的最大子数组和。
	for i := 1; i <= n; i++ {
		dp[i] = maxKConcatenationMaxSum(dp[i-1]+arr[i-1], arr[i-1])
		maxSum = maxKConcatenationMaxSum(maxSum, dp[i])
	}

	demo := 0
	// 查看中间的一段是否可以加上
	if k > 2 && sum > 0 {
		demo = sum * (k - 2) % mod
	}

	return (demo + maxSum%mod) % mod
}

func maxKConcatenationMaxSum(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
