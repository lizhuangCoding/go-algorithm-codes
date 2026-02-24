package main

// 二、差分 2.1一维差分（扫描线）
// 力扣：https://leetcode.cn/problems/points-that-intersect-with-cars/description/

func numberOfPoints(nums [][]int) int {
	maxLen := 0
	for _, v := range nums {
		if maxLen < v[1] {
			maxLen = v[1]
		}
	}

	diff := make([]int, maxLen+10)
	for i := 0; i < len(nums); i++ {
		diff[nums[i][0]]++
		diff[nums[i][1]+1]--
		// 当然也可以不用差分数组，直接从left ++ 到right，然后看这个位置是否是0，如果是0，说明没有，如果不是0，说明有
	}

	sum := 0
	demo := 0
	for i := 0; i < len(diff); i++ {
		demo += diff[i]
		// 如果不是0，说明包含了
		if demo != 0 {
			sum++
		}
	}
	return sum
}
