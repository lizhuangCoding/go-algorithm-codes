package main

// 动态规划：01背包
// 力扣：https://leetcode.cn/problems/partition-equal-subset-sum/

// 方法一：二维
// func canPartition(nums []int) bool {
// 	sum := 0
// 	for _, v := range nums {
// 		sum += v
// 	}
//
// 	// 剪枝：不是偶数，直接退出
// 	if sum&1 == 1 {
// 		return false
// 	}
//
// 	n := len(nums)
//
// 	dp := make([][]int, n+1)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = make([]int, sum/2+1)
// 	}
//
// 	for i := 1; i <= n; i++ {
// 		// 倒序，正序都可以
// 		for j := 0; j <= sum/2; j++ {
//
// 			// 容量不足
// 			if j < nums[i-1] {
// 				dp[i][j] = dp[i-1][j]
// 			} else {
// 				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i-1]]+nums[i-1])
// 			}
//
// 		}
// 	}
//
// 	return dp[n][sum/2] == sum/2
// }

// 方法二：一维
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	// 剪枝：不是偶数，直接退出
	if sum&1 == 1 {
		return false
	}

	n := len(nums)

	dp := make([]int, sum/2+1)

	for i := 1; i <= n; i++ {
		// 必须倒序
		for j := sum / 2; j >= 0; j-- {
			// 容量不足
			if j < nums[i-1] {
				dp[j] = dp[j]
			} else {
				dp[j] = max(dp[j], dp[j-nums[i-1]]+nums[i-1])
			}
		}
	}

	return dp[sum/2] == sum/2
}

// func main() {
// 	fmt.Println(canPartition([]int{1, 5, 11, 5}))
// 	fmt.Println(canPartition([]int{1, 2, 3, 5}))
//
// }
