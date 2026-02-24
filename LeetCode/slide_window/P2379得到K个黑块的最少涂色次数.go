package main

import (
	"math"
)

// 滑动窗口
// 力扣：https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/description/

func minimumRecolors(blocks string, k int) int {
	wNum := 0
	minNeed := math.MaxInt64
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			wNum++
		}
		if i >= k-1 {
			if wNum < minNeed {
				minNeed = wNum
			}
			if blocks[i-(k-1)] == 'W' {
				wNum--
			}
		}
	}
	return minNeed
}
