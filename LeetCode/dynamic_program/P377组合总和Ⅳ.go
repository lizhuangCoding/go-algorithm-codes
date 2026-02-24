// package main
//
// // 动态规划 一、入门DP 1.1 爬楼梯
// // 力扣：https://leetcode.cn/problems/combination-sum-iv/description/
//
// // 本质是爬楼梯，相当于每次往上爬 nums[i] 步
// func combinationSum4(nums []int, target int) int {
// 	dp := make([]int, target+1)
// 	dp[0] = 1
//
// 	for i := 1; i <= target; i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if nums[j] <= i {
// 				dp[i] += dp[i-nums[j]]
// 			}
// 		}
//
// 		// nums = [1,2,3], target = 4
// 		// 输出：7
// 		// dp[1] = dp[0]
// 		// dp[2] = dp[1] + dp[0]
// 		// dp[3] = dp[2] + dp[1] + dp[0]
//
// 	}
// 	return dp[target]
// }

package main

// 动态规划：完全背包的排列问题
// 力扣：https://leetcode.cn/problems/combination-sum-iv/description/
// 注意：顺序不同的序列被视作不同的组合。数组中的元素不重复。

// 二维dp
// func combinationSum4(nums []int, target int) int {
// 	// dp[i][j]: 用前i个数字凑出j的排列数（i的实际意义是“数字的使用次数”）
// 	dp := make([][]int, target+1)
// 	for i := range dp {
// 		dp[i] = make([]int, len(nums)+1)
// 	}
//
// 	// 初始化：和为0的排列数为1（空集合）
// 	for i := 0; i <= len(nums); i++ {
// 		dp[0][i] = 1
// 	}
//
// 	// 排列问题，而不是组合问题，所以要先遍历容量，然后遍历物品
// 	for j := 1; j <= target; j++ { // 先遍历容量
// 		for i := 1; i <= len(nums); i++ { // 再遍历数字
//
// 			dp[j][i] = dp[j][i-1] // 不选nums[i-1]
//
// 			if j >= nums[i-1] {
// 				dp[j][i] += dp[j-nums[i-1]][len(nums)] // 选nums[i-1]，注意用全部数字
// 			}
//
// 		}
// 	}
//
// // 	for i := 0; i < len(dp); i++ {
// // 		fmt.Println(dp[i])
// // 	}
//
// 	return dp[target][len(nums)]
// }

// 一维dp
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1

	// 排列问题，而不是组合问题，所以要先遍历容量，然后遍历物品
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {

			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}

		}
	}

	// for i := 0; i < len(dp); i++ {
	// 	fmt.Println(dp[i])
	// }

	return dp[target]
}

// func main() {
// 	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
// }
