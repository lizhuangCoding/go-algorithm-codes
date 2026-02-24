package main

import "sort"

// 签到：数组

func numberGame(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] <= nums[j]
	})

	arr := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		n1, n2 := nums[i], nums[i+1]
		arr = append(arr, n2)
		arr = append(arr, n1)
	}
	return arr
}
