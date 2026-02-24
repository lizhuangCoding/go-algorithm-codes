package main

// 动态规划
// 力扣：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

// func maxProfit(prices []int) int {
// 	n := len(prices)
//
// 	dp := make([][2]int, n)
// 	// 0:不持有股票，1:持有股票
// 	dp[0][0] = 0
// 	dp[0][1] = -prices[0]
//
// 	for i := 1; i < n; i++ {
// 		// 当前没有股票，有两种情况：上一把没有买股票；把上一把的股票卖出
// 		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
//
// 		// 当前有股票，有两种情况：上一把买了股票，这把啥也不干；上一把没有买股票，这一把买股票（一次只能买一股股票）
// 		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
// 	}
//
// 	return dp[n-1][0]
// }
