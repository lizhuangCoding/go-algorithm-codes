package main

// 动态规划
// 力扣：https://leetcode.cn/problems/contiguous-sequence-lcci/

// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
//
// 	res := dp[0]
//
// 	for i := 1; i < len(nums); i++ {
// 		dp[i] = max(dp[i-1]+nums[i], nums[i])
// 		res = max(res, dp[i])
// 	}
//
// 	return res
// }

// 空间优化：
// func maxSubArray(nums []int) int {
// 	pre := nums[0]
// 	res := pre
//
// 	for i := 1; i < len(nums); i++ {
// 		now := max(pre+nums[i], nums[i])
// 		pre = now
// 		res = max(res, now)
// 	}
//
// 	return res
// }
