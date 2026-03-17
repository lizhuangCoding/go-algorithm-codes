package main

import (
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/find-polygon-with-the-largest-perimeter/description/

// // largestPerimeter 函数的目标是从给定的整数切片 nums 中选取一些元素作为多边形的边长，
// // 尝试找出能构成多边形的最大周长，如果无法构成多边形则返回 -1。
// func largestPerimeter(nums []int) int64 {
// 	// 使用 slices.Sort 函数对输入的切片 nums 进行升序排序。
// 	// 排序的目的是为了后续能从大到小处理这些边长，利用贪心策略，优先尝试用较长的边来构成周长最大的多边形。
// 	slices.Sort(nums)
//
// 	// 初始化变量 s 用于存储所有边长的总和。
// 	s := 0
// 	// 通过 for 循环遍历切片 nums 中的每个元素 x。
// 	for _, x := range nums {
// 		// 把每个元素 x 累加到总和 s 中。
// 		s += x
// 	}
//
// 	// 开启一个从切片末尾开始往前的循环，i 从切片长度减 1 开始，递减到 2 为止。
// 	// 因为要构成多边形至少需要三条边，所以 i 要大于 1。
// 	for i := len(nums) - 1; i > 1; i-- {
// 		// 取出当前最大的边长，赋值给变量 x。
// 		x := nums[i]
//
// 		// 判断除当前最大边 x 之外的其余边长度总和（即 s - x）是否大于当前最大边 x。
// 		// 这是基于多边形的一个必要条件：任意一条边的长度要小于其余所有边的长度总和。
// 		// s > x * 2 等同于 s - x > x。
// 		if s > x*2 {
// 			// 如果满足条件，说明当前选取的这些边可以构成一个多边形，
// 			// 直接将所有边的总和 s 转换为 int64 类型并返回，因为题目要求返回的类型是 int64。
// 			return int64(s)
// 		}
//
// 		// 如果当前最大边不满足构成多边形的条件，
// 		// 就把当前最大边的长度从总和 s 中减去，
// 		// 然后继续考虑下一个次大的边。
// 		s -= x
// 	}
//
// 	// 如果遍历完所有可能的组合都没有找到能构成多边形的边长组合，
// 	// 则返回 -1 表示无法构成多边形。
// 	return -1
// }

func largestPerimeter(nums []int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for i := 0; i < len(nums)-2; i++ { // i < len(nums)-2：必须要保证至少有三条边
		if nums[i] < sum-nums[i] {
			return int64(sum)
		}
		sum -= nums[i]
	}

	return -1
}
