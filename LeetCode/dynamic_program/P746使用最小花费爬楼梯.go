package main

// 动态规划 一、入门DP 1.1 爬楼梯
// 力扣：https://leetcode.cn/problems/min-cost-climbing-stairs/description/

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1) // 表示爬到第i个楼梯需要支付的费用
	dp[0] = 0
	dp[1] = 0 // 起点就是索引1，不用支付

	for i := 2; i <= n; i++ {
		if dp[i-1]+cost[i-1] <= dp[i-2]+cost[i-2] {
			dp[i] = dp[i-1] + cost[i-1]
		} else {
			dp[i] = dp[i-2] + cost[i-2]
		}
	}

	return dp[n]
}

// func main() {
// 	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                         // 15
// 	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) // 6
// }

// 或者：
// func minCostClimbingStairs(cost []int) int {
// 	dp := make([]int, len(cost))
// 	dp[0] = cost[0]
// 	dp[1] = cost[1]
//
// 	for i := 2; i < len(cost); i++ {
// 		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
// 	}
//
// 	return min(dp[len(cost)-1], dp[len(cost)-2])
// }
