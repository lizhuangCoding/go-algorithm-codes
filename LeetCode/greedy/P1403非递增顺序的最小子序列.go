package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/minimum-subsequence-in-non-increasing-order/description/

func minSubsequence(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	res := make([]int, 0)
	demo := 0
	for i := 0; i < len(nums); i++ {
		demo += nums[i]
		res = append(res, nums[i])

		if demo > sum-demo {
			break
		}
	}

	return res
}
