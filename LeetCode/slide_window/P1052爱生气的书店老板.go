package main

// 滑动窗口，不错的题目（一开始看成了动态规划）
// 力扣：https://leetcode.cn/problems/grumpy-bookstore-owner/description/

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	total := 0 // 满意的总人数
	for i := 0; i < len(grumpy); i++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	maxNum := 0 // 满意人数的最大结果
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 { // 如果不满意就改为满意
			total += customers[i]
		}
	}
	maxNum = total

	for i := minutes; i < len(grumpy); i++ {
		left := i - minutes
		// 减去不满意的人
		if grumpy[left] == 1 {
			total -= customers[left]
		}

		// 因为已经遍历过满意的人，所以只添加不满意的
		if grumpy[i] == 1 {
			total += customers[i]
		}
		if maxNum < total {
			maxNum = total
		}
	}
	return maxNum
}

// 或者：总满意人数+最大新增满意人数
// func maxSatisfied(customers []int, grumpy []int, minutes int) int {
// 	total := 0
// 	n := len(customers)
// 	for i := 0; i < n; i++ {
// 		if grumpy[i] == 0 {
// 			total += customers[i]
// 		}
// 	}
//
// 	increase := 0
// 	for i := 0; i < minutes; i++ {
// 		increase += customers[i] * grumpy[i]
// 	}
// 	maxIncrease := increase
// 	for i := minutes; i < n; i++ {
// 		increase = increase - customers[i-minutes]*grumpy[i-minutes] + customers[i]*grumpy[i]
// 		if maxIncrease < increase {
// 			maxIncrease = increase
// 		}
// 	}
// 	return total + maxIncrease
// }
