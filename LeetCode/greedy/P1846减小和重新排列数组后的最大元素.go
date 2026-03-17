package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-element-after-decreasing-and-rearranging/

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		arr[i] = min(arr[i], arr[i-1]+1)
	}

	return arr[len(arr)-1]
}
