package main

import (
	"math"
	"sort"
)

// 签到：排序

func minimumAverage(nums []int) float64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := math.MaxFloat64
	n := len(nums)
	for i := 0; i < n/2; i++ {
		var demo = (float64(nums[i]) + float64(nums[n-i-1])) / 2.0
		if res > demo {
			res = demo
		}
	}
	return res
}
