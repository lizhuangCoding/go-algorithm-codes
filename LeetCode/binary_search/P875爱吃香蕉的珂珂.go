package main

import "math"

// 二分查找
// 力扣：https://leetcode.cn/problems/koko-eating-bananas/

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 0
	result := math.MaxInt64

	for i := 0; i < len(piles); i++ {
		right = max(right, piles[i])
	}

	// 找到最左边界
	for left <= right {
		// 速度k
		mid := (right + left) / 2

		// 计算是否可以在规定时间内吃完
		time := 0
		for i := 0; i < len(piles); i++ {
			time += piles[i] / mid
			if piles[i]%mid != 0 {
				time++
			}
		}

		if time <= h {
			result = min(result, mid)
			right = mid - 1
		} else if time > h {
			left = mid + 1
		}
	}

	return result
}
