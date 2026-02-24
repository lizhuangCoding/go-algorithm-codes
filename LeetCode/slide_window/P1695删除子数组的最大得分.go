package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/maximum-erasure-value/description/

func maximumUniqueSubarray(nums []int) int {
	m := make(map[int]bool)
	maxNum, left := 0, 0
	demo := 0
	for right := 0; right < len(nums); right++ {
		for left < right && m[nums[right]] {
			delete(m, nums[left])
			demo -= nums[left]
			left++
		}

		m[nums[right]] = true
		demo += nums[right]

		if maxNum < demo {
			maxNum = demo
		}
	}

	return maxNum
}
