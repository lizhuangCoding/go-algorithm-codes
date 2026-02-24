package main

// 动态规划：完全背包+全排列
// 卡码网：https://kamacoder.com/problempage.php?pid=1067

// func main() {
// 	// 容量为n，物品数量为m
// 	n, m := 0, 0
// 	fmt.Scanf("%d %d", &n, &m)
// 	fmt.Println(climbStairs(n, m))
// }
//
// // 一维dp
// func climbStairs(n int, m int) int {
// 	dp := make([]int, n+1)
// 	dp[0] = 1
//
// 	// 全排列问题，所以要先遍历容量，然后遍历物品
// 	for i := 0; i <= n; i++ {
// 		for j := 1; j <= m; j++ {
// 			if i >= j {
// 				dp[i] += dp[i-j]
// 			}
// 		}
// 	}
// 	return dp[n]
// }

// // 二维dp：没怎么看明白
// func climbStairs(n int, m int) int {
// 	// 创建二维 DP 数组，dp[i][j] 表示到达第 i 阶且最后一步爬 j 个台阶的方法数
// 	dp := make([][]int, n+1)
// 	for i := range dp {
// 		dp[i] = make([]int, m+1)
// 	}
//
// 	// 正确初始化起点状态：可以从 0 阶直接爬 j 步到达 j 阶（j ≤ m）
// 	for j := 1; j <= m && j <= n; j++ {
// 		dp[j][j] = 1
// 	}
//
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= m && j <= i; j++ {
// 			if i == j {
// 				continue // 已初始化的情况跳过
// 			}
// 			// 累加所有可能的前一个状态
// 			for k := 1; k <= m; k++ {
// 				if i-j >= 0 {
// 					dp[i][j] += dp[i-j][k]
// 				}
// 			}
// 		}
// 	}
//
// 	// 计算最终结果：到达第 n 阶的所有可能方法数之和
// 	result := 0
// 	for j := 1; j <= m; j++ {
// 		result += dp[n][j]
// 	}
// 	return result
// }
