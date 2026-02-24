package main

// 动态规划 一、入门DP 1.2 打家劫舍
// 力扣：https://leetcode.cn/problems/delete-and-earn/description/

// 题解：
// 看思路把这道题转换为打家劫舍：https://leetcode.cn/problems/delete-and-earn/solutions/758491/zhe-xiao-tou-you-lai-qiang-jie-liao-ta-z-w29x/
// 看代码：https://leetcode.cn/problems/delete-and-earn/solutions/758416/shan-chu-bing-huo-de-dian-shu-by-leetcod-x1pu/

func deleteAndEarn(nums []int) int {
	maxRes := 0
	for _, v := range nums {
		maxRes = maxDeleteAndEarn(maxRes, v)
	}

	// 把 2,2,3,3,3,4 变为 0:0, 1:0, 2:4, 3:9, 4:4
	sum := make([]int, maxRes+1)
	for _, v := range nums {
		sum[v] += v
	}
	return deleteAndEarnRob(sum)
}

// 类似于打家劫舍：间隔偷东西
func deleteAndEarnRob(nums []int) int {
	// 根据题意，nums的长度一定>=2
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = maxDeleteAndEarn(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = maxDeleteAndEarn(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func maxDeleteAndEarn(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
