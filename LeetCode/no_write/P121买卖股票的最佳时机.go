package no_write

// 动态规划 五、状态机 DP
// 一般定义 f[i][j] 表示前缀 a[:i] 在状态j下的最优值。一般j都很小。代表题目是「买卖股票」系列。
// 力扣：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

// func maxProfit(prices []int) int {
// 	n := len(prices)
// 	dp := make([][2]int, n) // 0：
// 	dp[0][0] =
// }
//
// func maxProfit(prices []int) (ans int) {
// 	minPrice := prices[0]
// 	for _, p := range prices {
// 		ans = max(ans, p-minPrice)
// 		minPrice = min(minPrice, p)
// 	}
// 	return
// }
//
