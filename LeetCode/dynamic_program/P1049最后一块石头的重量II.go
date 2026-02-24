package main

// 动态规划：01背包
// 力扣：https://leetcode.cn/problems/last-stone-weight-ii/

// 方法一：二维dp
func lastStoneWeightII(stones []int) int {
	sum := 0
	n := len(stones)

	for _, v := range stones {
		sum += v
	}

	// 分为两堆：找到能填充容量为 sum/2 的石头重量的最大值
	target := sum / 2

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}

	for i := 1; i <= n; i++ { // 物品
		for j := 0; j <= target; j++ { // 容量

			// 容量不足
			if j < stones[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i-1]]+stones[i-1])
			}

		}
	}

	return sum - dp[n][target] - dp[n][target]
}

// 方法二：一维dp
// 略，参考 416分割等和子集

// func main() {
// 	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))   // 1
// 	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40})) // 5
// }
