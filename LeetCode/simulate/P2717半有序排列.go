package main

// 数组，模拟
// 力扣：https://leetcode.cn/problems/semi-ordered-permutation/description/

func semiOrderedPermutation(nums []int) int {
	index1, index2 := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			index1 = i
		}
		if nums[i] == len(nums) {
			index2 = i
		}
	}

	total := 0

	if index1 < index2 { // 1 在 n 的左边
		total = index1 + (len(nums) - 1 - index2)
	} else if index1 > index2 { // 1 在 n 的右边
		total = index1 + (len(nums) - 1 - index2) - 1 // 可以少移动一次
	}
	return total
}
