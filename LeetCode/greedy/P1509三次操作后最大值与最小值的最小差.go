package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/description/

// 排序后，四种情况：
/*

输入：nums = [1,5,0,10,14]
输出：1
解释：我们最多可以走 3 步。
第一步，将 5 改为 0 。 nums变成 [1,0,0,10,14] 。
第二步，将 10 改为 0 。 nums变成 [1,0,0,0,14] 。
第三步，将 14 改为 1 。 nums变成 [1,0,0,0,1] 。
执行 3 步后，最小值和最大值之间的差值为 1 - 0 = 1 。
可以看出，没有办法可以在 3 步内使差值变为0。

0 1 5 10 14：（让两侧的元素尽可能的向中间靠拢）
10 10 10 10 14
5 5 5 10 10
1 1 5 5 5
0 1 1 1 1

*/

func minDifference(nums []int) int {
	if nums == nil {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	n := len(nums)
	if n <= 4 {
		return 0
	}
	return min(nums[n-1]-nums[3], nums[n-2]-nums[2], nums[n-3]-nums[1], nums[n-4]-nums[0])
}
