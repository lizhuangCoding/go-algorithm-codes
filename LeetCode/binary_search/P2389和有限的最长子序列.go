package main

import "slices"

// 二分查找
// 力扣：https://leetcode.cn/problems/longest-subsequence-with-limited-sum/description/

func answerQueries(nums []int, queries []int) []int {
	// 这里的子序列指的是可以随意从nums中挑选值，所以可以对nums进行排列
	slices.Sort(nums)

	sli := make([]int, len(nums))
	sli[0] = nums[0]
	for i := 1; i < len(sli); i++ {
		sli[i] = sli[i-1] + nums[i]
	}

	res := make([]int, 0)

	for i := 0; i < len(queries); i++ {
		target := queries[i]

		index := -1
		// 找到 <= target 的最大索引位置的值
		l, r := 0, len(nums)-1
		for l <= r {
			mid := (l + r) >> 1
			if sli[mid] <= target {
				index = mid
				l = mid + 1
			} else if sli[mid] > target {
				r = mid - 1
			}
		}

		res = append(res, index+1)
	}

	return res
}
