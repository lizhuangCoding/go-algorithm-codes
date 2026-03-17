package main

import "math"

// 贪心
// 力扣：https://leetcode.cn/problems/maximum-distance-in-arrays/description/

// 遍历每个数组时，用当前数组的最小值和最大值，去和之前遇到的全局最小和全局最大比较，但不能是同一个数组
func maxDistance(arrays [][]int) int {
	minRes := math.MaxInt64 / 2
	maxRes := math.MinInt64 / 2

	res := 0
	for _, arr := range arrays {
		x, y := arr[0], arr[len(arr)-1] // 本轮的最大和最小值

		// 关键步骤：用当前数组和之前遍历过的所有数组计算最大距离，避免最大最小值都取的是同一个数组
		// y - mn: 当前数组最大值 - 之前的最小值
		// mx - x: 之前的最大值 - 当前数组最小值
		res = max(res, y-minRes, maxRes-x)

		// 更新全局最小和全局最大
		minRes = min(minRes, x)
		maxRes = max(maxRes, y)
	}

	return res
}

// 错误的写法：没有考虑到最大、最小值都取的是同一个数组
// func maxDistance(arrays [][]int) int {
// 	minRes := math.MaxInt64
// 	maxRes := math.MinInt64
//
// 	for i := 0; i < len(arrays); i++ {
// 		for j := 0; j < len(arrays[i]); j++ {
// 			minRes = min(minRes, arrays[i][0])
// 			maxRes = max(maxRes, arrays[i][len(arrays[i])-1])
// 		}
// 	}
// 	return maxRes - minRes
// }
