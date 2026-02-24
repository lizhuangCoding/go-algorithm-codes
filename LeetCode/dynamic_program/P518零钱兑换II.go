package main

// 动态规划：完全背包
// 力扣：https://leetcode.cn/problems/coin-change-ii/description/

// 二维dp：
func change(amount int, coins []int) int {
	// dp[i][j] 表示前 i 个数中，和为 j 的个数
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}

	// 初始化
	// 前 0 个数中，和为 0 的个数 为 0
	dp[0][0] = 1

	for i := 1; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {

			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}

		}
	}

	return dp[len(coins)][amount]
}

// 一维dp：
// func change(amount int, coins []int) int {
// 	dp := make([]int, amount+1)
//
// 	dp[0] = 1
//
// 	for i := 1; i <= len(coins); i++ {
// 		for j := 0; j <= amount; j++ {
//
// 			// if j < coins[i-1] {
// 			// 	dp[j] = dp[j]
// 			// } else {
// 			// 	dp[j] = dp[j] + dp[j-coins[i-1]]
// 			// }
//
// 			// 可以简略为：
// 			if j >= coins[i-1] {
// 				dp[j] += dp[j-coins[i-1]]
// 			}
//
// 		}
// 	}
//
// 	return dp[amount]
// }

// func main() {
// 	fmt.Println(change(5, []int{1, 2, 5})) // 4
// }

/*

dp:

   0 1 2 3 4 5
0 [1 0 0 0 0 0]
1 [1 1 1 1 1 1]
2 [1 1 2 2 3 3]
3 [1 1 2 2 3 4]

*/
