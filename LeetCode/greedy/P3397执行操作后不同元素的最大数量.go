package main

import (
	"math"
	"slices"
)

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-number-of-distinct-elements-after-operations/description/

// func maxDistinctElements(nums []int, k int) int {
// 	if nums == nil || len(nums) == 0 {
// 		return 0
// 	}
//
// 	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
//
// 	m := make(map[int]struct{})
// 	nums[0] -= k
// 	m[nums[0]] = struct{}{}
//
// 	for i := 1; i < len(nums); i++ {
// 		demo := min(max(nums[i-1]+1, nums[i]-k), nums[i]+k)
//
// 		nums[i] = demo
// 		m[demo] = struct{}{}
// 	}
//
// 	return len(m)
// }

// 空间优化：
func maxDistinctElements(nums []int, k int) (ans int) {
	slices.Sort(nums)
	pre := math.MinInt // 记录每个人左边的人的位置
	for _, x := range nums {
		x = min(max(x-k, pre+1), x+k)

		if x > pre {
			ans++
			pre = x
		}

	}
	return
}
