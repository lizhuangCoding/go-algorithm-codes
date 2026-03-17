package main

import (
	"sort"
)

func minIncrementForUnique(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	// 1 1 2 2 3 7
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			nums[i]++
			res++
		} else if nums[i] < nums[i-1] {
			res += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}

	return res
}
