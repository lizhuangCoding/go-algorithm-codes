package main

import (
	"slices"
)

// 动态规划 一、入门DP 1.2 打家劫舍
// 力扣：https://leetcode.cn/problems/maximum-total-damage-with-spell-casting/description/

// 仿照 P740删除并获得点数：（不能全部通过，问题是：数据范围 1 <= power[i] <= 10^9，切片的容量会超出范围）
// func maximumTotalDamage(power []int) int64 {
// 	maxRes := 0
// 	for _, v := range power {
// 		maxRes = maxMaximumTotalDamage(maxRes, v)
// 	}
//
// 	sum := make([]int, maxRes+1)
// 	for _, v := range power {
// 		sum[v] += v
// 	}
// 	return int64(maximumTotalDamageRob(sum))
// }
//
// // 类似于打家劫舍：间隔偷东西
// func maximumTotalDamageRob(nums []int) int {
// 	n := len(nums)
// 	dp := make([]int, n)
// 	dp[0] = nums[0]
// 	dp[1] = maxMaximumTotalDamage(nums[0], nums[1])
// 	if n <= 2 {
// 		return dp[n-1]
// 	}
// 	dp[2] = maxMaximumTotalDamage(nums[2], dp[1])
//
// 	for i := 3; i < n; i++ {
// 		dp[i] = maxMaximumTotalDamage(dp[i-3]+nums[i], maxMaximumTotalDamage(dp[i-1], dp[i-2]))
// 	}
//
// 	return dp[n-1]
// }
//
// func maxMaximumTotalDamage(a, b int) int {
// 	if a >= b {
// 		return a
// 	}
// 	return b
// }

func maximumTotalDamage(power []int) int64 {
	cnt := map[int]int{}
	for _, x := range power {
		cnt[x]++
	}

	n := len(cnt)
	a := make([]int, 0, n)
	for x := range cnt {
		a = append(a, x)
	}

	slices.Sort(a)
	// a : 1, 6, 7
	dp := make([]int, n+1) // dp[i] 表示到伤害值为 a[i-1] 时能得到的最大伤害总和。
	j := 0                 // 初始化 j = 0 是为了在循环中用来指向能与当前伤害值 a[i] 不冲突的最近的伤害值（可以debug试一试）
	for i, x := range a {
		// 不能使用伤害为 power[i] - 2 ，power[i] - 1 ，power[i] + 1 或者 power[i] + 2 的咒语。
		for a[j] < x-2 { // 找到与 x 不冲突（差距大于 2）的最大下标。例如，当 x = 5 时，只会选择那些小于 5-2 = 3 的伤害值。
			j++
		}
		// dp[i]：表示不选择 x 伤害值的最大伤害和。
		// dp[j] + x * cnt[x]：表示选择 x 伤害值（和它出现的次数 cnt[x] 相乘），并加上不冲突的最大和。
		dp[i+1] = maxMaximumTotalDamage(dp[i], dp[j]+x*cnt[x])
	}
	return int64(dp[n])
}

func maxMaximumTotalDamage(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(maximumTotalDamage([]int{7, 1, 6, 6}))
// }
