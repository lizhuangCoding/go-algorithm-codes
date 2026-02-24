package main

// 动态规划 一、入门DP 1.1 爬楼梯
// 力扣：https://leetcode.cn/problems/count-number-of-texts/

// 本质是爬楼梯（不是太理解）
func countTexts(pressedKeys string) int {
	n := len(pressedKeys)
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		ch := pressedKeys[i-1]

		end := 3
		// 按 7、9 时是特殊情况
		if ch == '7' || ch == '9' {
			end = 4
		}

		for j := 1; j <= end && j <= i; j++ {
			index := i - j
			// 如果不是连续的点击同一个数字，就退出
			if ch != pressedKeys[index] {
				break
			}
			dp[i] = (dp[i] + dp[index]) % 1_000_000_007
		}
	}
	return dp[n]
}
