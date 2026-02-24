package main

import (
	"math"
)

// 动态规划 一、入门DP 1.3 最大子数组和（最大子段和）
// 力扣：https://leetcode.cn/problems/maximum-subarray/description/

// 方法一：动态规划
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = maxMaxSubArray(nums[i], dp[i-1]+nums[i])
	}

	maxRes := math.MinInt
	for i := 0; i < n; i++ {
		maxRes = maxMaxSubArray(maxRes, dp[i])
	}
	return maxRes
}

func maxMaxSubArray(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 或者：
// func maxSubArray(nums []int) (ans int) {
// 	ans = nums[0]
//
// 	demo := 0
// 	for _, v := range nums {
// 		demo += v
// 		ans = maxMaxSubArray(ans, demo)
// 		if demo < 0 {
// 			demo = 0
// 		}
// 	}
// 	return ans
// }
