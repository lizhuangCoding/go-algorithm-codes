package main

// 动态规划 一、入门DP 1.2 打家劫舍
// 力扣：https://leetcode.cn/problems/house-robber-ii/description/

// 这个地方所有的房屋都围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
// 要么从索引 0 -> 索引len(nums)-2；要么从索引 1 -> 索引len(nums)-1
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// nums[:len(nums)-1]， robSolve(nums[1:])最少有1个元素、2个元素、3个元素...
	return maxRob2(rob2Solve(nums[:n-1]), rob2Solve(nums[1:]))
}

func rob2Solve(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = maxRob2(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = maxRob2(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[n-1]
}

func maxRob2(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 空间优化：
func robDemo2(nums []int) int {
	first, second := nums[0], maxRob2(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, maxRob2(first+v, second)
	}
	return second
}
