package main

// 动态规划 二、网格图DP 2.1 基础
// 力扣：https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/

func jewelleryValue(frame [][]int) int {
	m := len(frame)
	n := len(frame[0])

	// 如果把dp的长度改为 m+1,n+1，就不用初始化了
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化
	dp[0][0] = frame[0][0]
	// 初始化第一列
	for i := 1; i < m; i++ {
		dp[i][0] += frame[i][0] + dp[i-1][0]
	}
	// 初始化第一行
	for j := 1; j < n; j++ {
		dp[0][j] += frame[0][j] + dp[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = maxJewelleryValue(dp[i-1][j], dp[i][j-1]) + frame[i][j]
		}
	}
	return dp[m-1][n-1]
}

func maxJewelleryValue(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 空间优化：
// 由于 f[i+1] 只依赖 f[i]，那么 f[i−1] 及其之前的数据就没用了。例如计算 f[2] 的时候，数组 f[0] 不再使用了。
// 那么干脆把 f[2] 填到 f[0] 中。同理，把 f[3] 填到 f[1] 中，f[4] 填到 f[0] 中，…… 因此可以只用两个长为 n+1 的数组滚动计算。
// 时间复杂度：O(mn)。空间复杂度：O(n)。
// func jewelleryValue(grid [][]int) int {
// 	m, n := len(grid), len(grid[0])
// 	f := [2][]int{make([]int, n+1), make([]int, n+1)}
// 	for i, row := range grid {
// 		for j, x := range row {
// 			f[(i+1)%2][j+1] = maxJewelleryValue(f[(i+1)%2][j], f[i%2][j+1]) + x
// 		}
// 	}
// 	return f[m%2][n]
// }
