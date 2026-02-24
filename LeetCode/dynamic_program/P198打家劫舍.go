package main

// 动态规划 一、入门DP 1.2 打家劫舍
// 力扣：https://leetcode.cn/problems/house-robber/description/

// 偷窃第 k 间房屋，那么就不能偷窃第 k−1 间房屋，偷窃总金额为前 k−2 间房屋的最高总金额与第 k 间房屋的金额之和。
// 不偷窃第 k 间房屋，偷窃总金额为前 k−1 间房屋的最高总金额。
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = maxRob(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = maxRob(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func maxRob(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
