package main

// 动态规划 一、入门DP 1.3 最大子数组和（最大子段和）
// 力扣：https://leetcode.cn/problems/maximum-absolute-sum-of-any-subarray/description/

// 方法一：动态规划
func maxAbsoluteSum(nums []int) int {
	fMax, fMin := 0, 0
	res := 0

	for i := 0; i < len(nums); i++ {
		fMax = maxMaxAbsoluteSum(fMax, 0) + nums[i]
		fMin = minMaxAbsoluteSum(fMin, 0) + nums[i]
		res = maxMaxAbsoluteSum(res, maxMaxAbsoluteSum(fMax, -fMin))
	}

	return res
}

func maxMaxAbsoluteSum(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minMaxAbsoluteSum(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 方法二：前缀和
// func maxAbsoluteSum(nums []int) int {
// 	var s, mx, mn int
//
// 	for i := 0; i < len(nums); i++ {
// 		s += nums[i] // 前缀和
// 		if mx < s {
// 			mx = s
// 		}
// 		if mn > s {
// 			mn = s
// 		}
// 	}
// 	return mx - mn
// }

// func main() {
// 	fmt.Println(maxAbsoluteSum([]int{1, -3, 2, 3, -4}))                                                          // 5
// 	fmt.Println(maxAbsoluteSum([]int{2, -5, 1, -4, 3, -2}))                                                      // 8
// 	fmt.Println(maxAbsoluteSum([]int{-7, -1, 0, -2, 1, 3, 8, -2, -6, -1, -10, -6, -6, 8, -4, -9, -4, 1, 4, -9})) // 44
// }
