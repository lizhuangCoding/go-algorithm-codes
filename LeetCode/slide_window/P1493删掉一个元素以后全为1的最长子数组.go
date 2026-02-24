package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/description/

// func longestSubarray(nums []int) int {
// 	num0, num1 := 0, 0 // 0和1的数量
// 	maxLen := 0
// 	left := 0
// 	for right := 0; right < len(nums); right++ {
// 		if nums[right] == 0 {
// 			num0++
// 		} else {
// 			num1++
// 		}
//
// 		for num0 >= 2 {
// 			if nums[left] == 0 {
// 				num0--
// 			} else {
// 				num1--
// 			}
// 			left++
// 		}
// 		if maxLen < num1 {
// 			maxLen = num1
// 		}
// 	}
//
// 	// 因为必须要删除一个元素，所以为了防止元素全部是1，所以--
// 	if maxLen == len(nums) {
// 		maxLen--
// 	}
// 	return maxLen
// }

// 优化：不用++，--，而是直接+nums[right]，或者-nums[left]
func longestSubarray(nums []int) int {
	sum, left, right := 0, 0, 0
	for ; right < len(nums); right++ {
		sum += nums[right]
		if sum < right-left {
			sum -= nums[left]
			left++
		}
	}
	// 直接包含了数组元素全是1的情况，和减去一个0的长度的情况
	return right - left - 1
}
