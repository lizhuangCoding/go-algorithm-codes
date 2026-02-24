package main

// 双指针（贪心）
// 力扣：https://leetcode.cn/problems/container-with-most-water/description/

// 方法一：暴力，超时
// func maxArea(height []int) int {
// 	result := 0
// 	for i := 0; i < len(height); i++ {
// 		for j := i; j < len(height); j++ {
// 			temp := (j - i) * minMaxArea(height[i], height[j])
// 			result = maxMaxArea(result, temp)
// 		}
// 	}
//
// 	return result
// }

func maxMaxArea(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minMaxArea(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 方法二：double_pointer
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0

	for left < right {
		temp := (right - left) * minMaxArea(height[left], height[right])
		result = maxMaxArea(result, temp)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
