package main

import (
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/find-subsequence-of-length-k-with-the-largest-sum/description/

func maxSubsequence(nums []int, k int) []int {
	arr := make([]int, 0)
	arr = append(arr, nums...)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	m := make(map[int]int)
	for i := 0; i < len(arr[:k]); i++ {
		m[arr[i]]++
	}

	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		v, ok := m[nums[i]]
		if ok && v > 0 {
			res = append(res, nums[i])
			m[nums[i]]--
		}
	}

	return res
}
