package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/maximum-erasure-value/description/

func maxSubarrayLength(nums []int, k int) int {
	m := make(map[int]int)
	maxLen := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		// 没有添加前 m[nums[right]]已经==k了，所以需要减去之前的一个 nums[right]才能继续添加 nums[right]
		for left < right && m[nums[right]] == k {
			m[nums[left]]--
			if m[nums[left]] == 0 {
				delete(m, nums[left])
			}
			left++
		}

		m[nums[right]]++

		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// 或者先 m[nums[right]]++，然后--
// func maxSubarrayLength(nums []int, k int) int {
// 	length := len(nums)
// 	m := make(map[int]int, length)
// 	ans := 0
//
// 	for left, right := 0, 0; right < length; right++ {
// 		m[nums[right]]++
// 		for m[nums[right]] > k {
// 			m[nums[left]]--
// 			left++
// 		}
//
// 		if ans < right-left+1 {
// 			ans = right - left + 1
// 		}
// 	}
// 	return ans
// }
