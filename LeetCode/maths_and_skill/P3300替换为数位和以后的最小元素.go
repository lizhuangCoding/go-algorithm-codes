package main

import "math"

// 签到：数学

func minElement(nums []int) int {
	res := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		demo := 0
		for nums[i] > 0 {
			demo += nums[i] % 10
			nums[i] /= 10
		}
		if res > demo {
			res = demo
		}
	}
	return res
}
