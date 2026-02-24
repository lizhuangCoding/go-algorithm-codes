package main

// 数组
// 力扣：https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slowIndex, fastIndex := 1, 2
	for ; fastIndex < len(nums); fastIndex++ {
		if nums[slowIndex-1] != nums[fastIndex] {
			slowIndex++
			nums[slowIndex] = nums[fastIndex]
		}
	}

	return slowIndex + 1
}
