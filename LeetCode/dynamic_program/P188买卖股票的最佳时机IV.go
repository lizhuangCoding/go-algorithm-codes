package main

// 动态规划
// 力扣：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/

// k次交易，2 * k种状态
// 状态从1开始计算，避免判断
// 奇数时持有(保持或买入)
// 偶数时不持有(保持或卖出)

// func maxProfit(k int, prices []int) int {
// 	// dp[i][j]：
// 	// i：表示第 i 天（prices[i] 是当天的股价）
// 	// j：表示交易状态，范围 [1, 2*k]（共 2*k 种状态）
// 	dp := make([][]int, len(prices))
//
// 	for i := 0; i < len(prices); i++ {
// 		dp[i] = make([]int, 2*k+1)
// 	}
//
// 	// 奇数状态时持有
// 	for i := 0; i < len(dp[0]); i++ {
// 		if i%2 == 1 {
// 			dp[0][i] = -prices[0]
// 		}
// 	}
//
// 	// 奇数状态（j % 2 == 1）：当前持有股票（可能是第 (j+1)/2 次交易的买入状态）
// 	// 偶数状态（j % 2 == 0）：当前不持有股票（可能是第 j/2 次交易的卖出状态）
// 	//
// 	// 状态转移：
// 	// 持有股票（奇数 j）：
// 	// 保持持有前一天的股票：dp[i-1][j]，当天买入：dp[i-1][j-1] - prices[i]
// 	// 不持有股票（偶数 j）：
// 	// 保持不持有：dp[i-1][j]，当天卖出：dp[i-1][j-1] + prices[i]
//
// 	for i := 1; i < len(prices); i++ {
// 		for j := 1; j < len(dp[0]); j++ {
// 			if j%2 == 1 { // 奇数：持有
// 				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
// 			} else { // 偶数：卖出
// 				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
// 			}
// 		}
// 	}
//
// 	return dp[len(prices)-1][2*k]
// }
