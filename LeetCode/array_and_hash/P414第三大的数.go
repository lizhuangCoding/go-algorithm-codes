package main

import "sort"

// 数组：排序
// 力扣：https://leetcode.cn/problems/third-maximum-number/description/

func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // 倒序排序
	for i, diff := 1, 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			diff++
			if diff == 3 { // 此时 nums[i] 就是第三大的数
				return nums[i]
			}
		}
	}
	return nums[0]
}
