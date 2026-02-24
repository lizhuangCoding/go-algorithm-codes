package main

// 动态规划 一、入门DP 1.3 最大子数组和（最大子段和）
// 力扣：https://leetcode.cn/problems/maximum-score-of-spliced-array/description/

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	return maxMaximumsSplicedArray(maximumsSplicedArraySolve(nums1, nums2), maximumsSplicedArraySolve(nums2, nums1))
}

func maximumsSplicedArraySolve(nums1 []int, nums2 []int) int {
	diff := 0    // 计算的是差值
	maxDiff := 0 // 差值的最大值

	sum := 0
	for i := 0; i < len(nums1); i++ {
		sum += nums1[i]
		diff = maxMaximumsSplicedArray(diff+nums2[i]-nums1[i], 0)
		maxDiff = maxMaximumsSplicedArray(maxDiff, diff)
	}
	// 返回总和+最大差值
	return sum + maxDiff
}

func maxMaximumsSplicedArray(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
