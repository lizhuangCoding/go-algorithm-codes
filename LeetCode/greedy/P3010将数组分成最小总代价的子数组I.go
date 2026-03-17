package main

import (
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-i/description/

// 第一个必选，然后从剩下的元素中选出两个最小的
func minimumCost(nums []int) int {
	res := nums[0]

	arr := nums[1:]
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	res += arr[0]
	res += arr[1]

	return res
}

// 或者：维护最小值和次小值
// func minimumCost(nums []int) int {
// 	fi, se := math.MaxInt, math.MaxInt
// 	for _, x := range nums[1:] {
// 		if x < fi {
// 			se = fi
// 			fi = x
// 		} else if x < se {
// 			se = x
// 		}
// 	}
// 	return nums[0] + fi + se
// }
