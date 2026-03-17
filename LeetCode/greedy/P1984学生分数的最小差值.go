package main

import (
	"math"
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/description/

func minimumDifference(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := math.MaxInt64
	for i := 0; i < len(nums)-k+1; i++ {
		res = min(res, nums[i+k-1]-nums[i])
	}

	return res
}
