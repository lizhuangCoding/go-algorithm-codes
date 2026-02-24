package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/max-consecutive-ones-iii/description/

func longestOnes(nums []int, k int) int {
	maxLen := 0
	left := 0
	num0, num1 := 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			num0++
		} else {
			num1++
		}

		for num0 > k {
			if nums[left] == 0 {
				num0--
			} else {
				num1--
			}
			left++
		}

		if maxLen < num0+num1 {
			maxLen = num0 + num1
		}
	}
	return maxLen
}
