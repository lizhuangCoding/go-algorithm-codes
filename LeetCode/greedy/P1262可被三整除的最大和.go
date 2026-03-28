package main

import (
	"sort"
)

// 贪心
// 力扣：https://leetcode.cn/problems/greatest-sum-divisible-by-three/description/

/*
1. 如果总和能被3整除，直接返回总和

2. 如果总和除以3余1，需要删除：
一个最小的余数为1的数，或
两个最小的余数为2的数

3. 如果总和除以3余2，需要删除：
一个最小的余数为2的数，或
两个最小的余数为1的数
*/
func maxSumDivThree(nums []int) int {
	sum := 0
	remainder1 := make([]int, 0) // 余数为1的数
	remainder2 := make([]int, 0) // 余数为2的数

	// 计算总和并分类
	for _, num := range nums {
		sum += num
		if num%3 == 1 {
			remainder1 = append(remainder1, num)
		} else if num%3 == 2 {
			remainder2 = append(remainder2, num)
		}
	}

	// 排序
	sort.Ints(remainder1)
	sort.Ints(remainder2)

	// 计算需要减去的值
	remainder := sum % 3

	if remainder == 0 {
		return sum
	}

	if remainder == 1 {
		// 方案1：减去一个最小的余数为1的数
		option1 := 0
		if len(remainder1) > 0 {
			option1 = sum - remainder1[0]
		}

		// 方案2：减去两个最小的余数为2的数
		option2 := 0
		if len(remainder2) >= 2 {
			option2 = sum - (remainder2[0] + remainder2[1])
		}

		return max(option1, option2)
	} else { // remainder == 2
		// 方案1：减去一个最小的余数为2的数
		option1 := 0
		if len(remainder2) > 0 {
			option1 = sum - remainder2[0]
		}

		// 方案2：减去两个最小的余数为1的数
		option2 := 0
		if len(remainder1) >= 2 {
			option2 = sum - (remainder1[0] + remainder1[1])
		}

		return max(option1, option2)
	}
}
