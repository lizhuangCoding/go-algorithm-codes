package main

import "math"

// 滑动窗口
// 力扣：https://leetcode.cn/problems/maximum-average-subarray-i/
func findMaxAverage(nums []int, k int) float64 {
	sum := math.MinInt
	demo := 0
	for i := 0; i <= k-1; i++ {
		demo += nums[i]
	}
	if demo > sum {
		sum = demo
	}

	for right := k; right < len(nums); right++ {
		left := right - k
		demo -= nums[left]
		demo += nums[right]

		if demo > sum {
			sum = demo
		}
	}
	return float64(sum) / float64(k)
}

// 代码简洁度优化:
// func findMaxAverage(nums []int, k int) float64 {
// 	s := 0
// 	mx := math.MinInt
// 	for i, v := range nums {
// 		s += v
// 		if i >= k {
// 			s -= nums[i-k]
// 		}
// 		if i >= k-1 {
// 			mx = max(mx, s)
// 		}
// 	}
// 	return float64(mx) / float64(k)
// }
