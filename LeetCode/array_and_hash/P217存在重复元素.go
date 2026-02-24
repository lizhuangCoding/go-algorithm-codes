package main

// 哈希表

// 签到
// 方法一：排序
// func containsDuplicate(nums []int) bool {
// 	mergeSort.Ints(nums)
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == nums[i-1] {
// 			return true
// 		}
// 	}
// 	return false
// }

// 方法二：哈希表
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, n := range nums {
		if m[n] {
			return true
		}
		m[n] = true
	}
	return false
}
