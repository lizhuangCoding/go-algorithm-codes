package main

// 数组
// 力扣：https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

// func removeDuplicates(nums []int) int {
// 	slowIndex := 0
//
// 	for fastIndex := 1; fastIndex < len(nums); fastIndex++ {
// 		if nums[slowIndex] != nums[fastIndex] {
// 			slowIndex++
// 			nums[slowIndex] = nums[fastIndex]
// 		}
// 	}
//
// 	return slowIndex + 1
// }
