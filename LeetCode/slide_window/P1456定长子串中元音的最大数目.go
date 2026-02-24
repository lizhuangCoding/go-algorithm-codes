package main

import (
	"strings"
)

// 滑动窗口
// 力扣：https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/

func maxVowels(s string, k int) int {
	if k <= 0 {
		return 0
	}

	r := "aeiou"
	num := 0
	// 这个叫双指针，不算滑动窗口，要注意：
	// for left, right := 0, k-1; left <= right && right < len(s); left, right = left+1, right+1 {
	// 	index := left
	// 	demo := 0
	// 	for index <= right {
	// 		if strings.Contains(r, string(s[index])) {
	// 			demo++
	// 		}
	// 		num = maxMaxVowels(num, demo)
	// 		index++
	// 	}
	// }

	// 优化为：
	demo := 0
	for i := 0; i <= k-1; i++ {
		if strings.Contains(r, string(s[i])) {
			demo++
		}
	}
	num = maxMaxVowels(num, demo)

	for right := k; right < len(s); right++ {
		left := right - k
		if strings.Contains(r, string(s[left])) {
			demo--
		}
		if strings.Contains(r, string(s[right])) {
			demo++
		}
		num = maxMaxVowels(num, demo)
	}

	return num
}

func maxMaxVowels(a, b int) int {
	if a > b {
		return a
	}
	return b
}
