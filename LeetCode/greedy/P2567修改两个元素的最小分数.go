package main

import "sort"

// 贪心
// 力扣：https://leetcode.cn/problems/minimum-score-by-changing-two-elements/description/

/*

排序后三种情况：
题解：https://leetcode.cn/problems/minimum-score-by-changing-two-elements/solutions/2119454/nao-jin-ji-zhuan-wan-by-endlesscheng-9l4m/

1 4 5 7 8：
三种情况：
1 4 5 5 5
4 4 5 7 7
5 5 5 7 8

*/

func minimizeSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	n := len(nums)
	return min(nums[n-3]-nums[0], nums[n-1]-nums[2], nums[n-2]-nums[1])
}
