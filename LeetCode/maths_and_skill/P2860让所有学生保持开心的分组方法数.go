package main

import "sort"

// 脑筋急转弯
// 力扣：https://leetcode.cn/problems/happy-students/

func countWays(nums []int) int {
	n := len(nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := 0

	// 全部学生都不选
	if nums[0] > 0 {
		res++
	}

	for i := 0; i < n-1; i++ {
		if i+1 > nums[i] && i+1 < nums[i+1] {
			res++
		}
	}

	// 选择所有学生
	if n > nums[n-1] {
		res++
	}

	return res
}
