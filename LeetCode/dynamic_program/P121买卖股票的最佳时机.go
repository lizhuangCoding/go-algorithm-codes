package main

import "math"

// 动态规划
// 力扣：https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	res := 0
	minPrice := math.MaxInt64 // 记录低点

	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}

		// 如果可以盈利，就卖出，记录盈利的值
		if prices[i]-minPrice > res {
			res = prices[i] - minPrice
		}
	}
	return res
}
