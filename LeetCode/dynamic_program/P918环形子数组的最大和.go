package main

import (
	"math"
)

// 动态规划 一、入门DP 1.3 最大子数组和（最大子段和）
// 力扣：https://leetcode.cn/problems/maximum-sum-circular-subarray/

// 分类讨论一:子数组没有跨过边界(在 nums 中间),此时算的是最大非空子数组和 (力扣第 53 题)，记作 maxS.
// 分类讨论二:子数组跨过边界(在 nums 两端)，由于子数组和+其余元素和 = sum(nums) = 常数，所以其余元素和越小，子数组和越大。
func maxSubarraySumCircular(nums []int) int {
	maxSum := math.MinInt // 最大子数组和，不能为空
	minSum := 0           // 最小子数组和，可以为空
	sum := 0              // 元素的总和
	maxF, minF := 0, 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		maxF = maxMaxSubarraySumCircular(maxF+nums[i], nums[i])
		maxSum = maxMaxSubarraySumCircular(maxSum, maxF)

		minF = minMaxSubarraySumCircular(minF+nums[i], nums[i])
		minSum = minMaxSubarraySumCircular(minSum, minF)
	}

	// 防止例子输出sum-minSum=0，应该输出-2：-3,-2,-3
	if sum == minSum {
		return maxSum
	}
	return maxMaxSubarraySumCircular(maxSum, sum-minSum)
}

func maxMaxSubarraySumCircular(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minMaxSubarraySumCircular(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// // 或者：力扣官方（没具体看）
// func maxSubarraySumCircular(nums []int) int {
// 	n := len(nums)
// 	leftMax := make([]int, n)
// 	// 对坐标为 0 处的元素单独处理，避免考虑子数组为空的情况
// 	leftMax[0] = nums[0]
// 	leftSum, pre, res := nums[0], nums[0], nums[0]
// 	for i := 1; i < n; i++ {
// 		pre = maxMaxSubarraySumCircular(pre+nums[i], nums[i])
// 		res = maxMaxSubarraySumCircular(res, pre)
// 		leftSum += nums[i]
// 		leftMax[i] = maxMaxSubarraySumCircular(leftMax[i-1], leftSum)
// 	}
// 	// 从右到左枚举后缀，固定后缀，选择最大前缀
// 	rightSum := 0
// 	for i := n - 1; i > 0; i-- {
// 		rightSum += nums[i]
// 		res = maxMaxSubarraySumCircular(res, rightSum+leftMax[i-1])
// 	}
// 	return res
// }

// func main() {
// 	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}))     // 10
// 	fmt.Println(maxSubarraySumCircular([]int{3, -2, 2, -3})) // 3
// 	fmt.Println(maxSubarraySumCircular([]int{-3, -2, -3}))   // -2
// }
