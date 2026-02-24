package main

import "slices"

// 二分查找
// 力扣：https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/

func successfulPairs(spells []int, potions []int, success int64) []int {
	res := make([]int, 0)

	slices.Sort(potions)

	for i := 0; i < len(spells); i++ {

		left, right := 0, len(potions)-1
		index := len(potions)

		// 找到最左边符合条件的值
		for left <= right {
			mid := left + (right-left)>>1

			if potions[mid]*spells[i] >= int(success) {
				index = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		res = append(res, len(potions)-index)
	}

	return res
}
