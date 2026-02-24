package main

import (
	"sort"
)

// 数学：排序
// 力扣：https://leetcode.cn/problems/maximum-product-of-three-numbers/description/

// 方法一：排序
// 先排序。
// 如果数组中全是非负数，则排序后最大的三个数相乘即为最大乘积；如果全是非正数，则最大的三个数相乘同样也为最大乘积。
// 如果数组中有正数有负数，则最大乘积既可能是三个最大正数的乘积，也可能是两个最小负数与最大正数的乘积。
// 综上，我们在给数组排序后，分别求出三个最大正数的乘积，以及两个最小负数与最大正数的乘积，二者之间的最大值即为所求答案。
// 时间复杂度：O(NlogN)
// 空间复杂度：O(logN)
func maximumProduct(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	return maxMaximumProduct(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}

func maxMaximumProduct(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法二：线性扫描
// 在方法一中，我们实际上只要求出数组中最大的三个数以及最小的两个数，因此我们可以不用排序，用线性扫描直接得出这五个数。
// func maximumProduct(nums []int) int {
// 	// 最小的和第二小的
// 	min1, min2 := math.MaxInt64, math.MaxInt64
// 	// 最大的、第二大的和第三大的
// 	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
//
// 	for _, x := range nums {
// 		if x < min1 {
// 			min2 = min1
// 			min1 = x
// 		} else if x < min2 {
// 			min2 = x
// 		}
//
// 		if x > max1 {
// 			max3 = max2
// 			max2 = max1
// 			max1 = x
// 		} else if x > max2 {
// 			max3 = max2
// 			max2 = x
// 		} else if x > max3 {
// 			max3 = x
// 		}
// 	}
// 	return maxMaximumProduct(min1*min2*max1, max1*max2*max3)
// }
