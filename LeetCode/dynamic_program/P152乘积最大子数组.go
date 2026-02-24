package main

// 动态规划 一、入门DP 1.3 最大子数组和（最大子段和）
// 力扣：https://leetcode.cn/problems/maximum-product-subarray/

// func maxProduct(nums []int) int {
// 	n := len(nums)
//  // maxDp[i] 和 minDp[i]：分别是以第 i 个元素结尾的子数组的最大和最小乘积。
// 	maxDp := make([]int, n)
// 	minDp := make([]int, n)
//
// 	// 初始化
// 	maxDp[0] = nums[0]
// 	minDp[0] = nums[0]
//
// 	for i := 1; i < n; i++ {
// 		maxDp[i] = maxMaxProduct(maxDp[i-1]*nums[i], maxMaxProduct(nums[i], minDp[i-1]*(nums[i])))
// 		minDp[i] = minMaxProduct(minDp[i-1]*nums[i], minMaxProduct(nums[i], maxDp[i-1]*(nums[i])))
//
// 		// 不写也可以：
// 		if minDp[i] < -1<<31 {
// 			minDp[i] = nums[i]
// 		}
// 	}
//
// 	res := maxDp[0]
// 	for i := 1; i < n; i++ {
// 		res = maxMaxProduct(res, maxDp[i])
// 	}
// 	return res
// }

// 优化空间：
// 正负数问题：正数乘以正数越乘越大，但负数乘以负数却可能变为正数，因此需要同时追踪最大值和最小值。
func maxProduct(nums []int) int {
	// maxF: 记录以当前元素结尾的子数组的最大乘积。 minF: 记录以当前元素结尾的子数组的最小乘积（因为负数可能通过乘以另一个负数变成最大）。
	maxF, minF, ans := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		// maxDemo 和 minDemo 保存上一轮的 maxF 和 minF
		maxDemo, minDemo := maxF, minF
		maxF = maxMaxProduct(maxDemo*nums[i], maxMaxProduct(nums[i], minDemo*nums[i]))
		minF = minMaxProduct(minDemo*nums[i], minMaxProduct(nums[i], maxDemo*nums[i]))

		// 如果 minF 超过整型最小值（如题解中 minF < -1 << 31），将其重置为当前值 nums[i]，避免无意义的大值。（不写这行代码似乎也可以）
		if minF < (-1 << 31) {
			minF = nums[i]
		}

		ans = maxMaxProduct(maxF, ans)
	}
	return ans
}

func maxMaxProduct(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minMaxProduct(x, y int) int {
	if x < y {
		return x
	}
	return y
}
