package main

import (
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-alternating-sum-of-squares/

func maxAlternatingSum(nums []int) int64 {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	res := 0
	for i := 0; i < len(nums); i++ {
		if i < (len(nums)+1)/2 {
			res += nums[i]
		} else {
			res -= nums[i]
		}
	}

	return int64(res)
}
