package main

// 数组，哈希表
// 力扣：https://leetcode.cn/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/

// 暴力
func minimumOperations(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	index := -1
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] >= 2 {
			index = i    // 应该截取到这里
			m[nums[i]]-- // 减少一个元素
		}
	}

	num := (index + 1) / 3
	if (index+1)%3 != 0 {
		num++
	}

	return num
}

// 倒序遍历
// func minimumOperations(nums []int) int {
// 	seen := make([]bool, 128)
//
// 	for i := len(nums) - 1; i >= 0; i-- {
// 		// 如果重复了，就返回
// 		if seen[nums[i]] {
// 			return i/3 + 1
// 		}
// 		seen[nums[i]] = true
// 	}
//
// 	return 0
// }
