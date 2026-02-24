package main

import (
	"strings"
)

// 字符串
// 力扣：https://leetcode.cn/problems/valid-palindrome/description/

// 方法一：筛选 + 判断
// func isPalindrome(s string) bool {
// 	s = strings.ToLower(s)
// 	newStr := ""
// 	for i := 0; i < len(s); i++ {
// 		if (s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57) {
// 			newStr = newStr + string(s[i])
// 		}
// 	}
//
//  // double_pointer
// 	left, right := 0, len(newStr)-1
// 	for left < right {
// 		if newStr[left] != newStr[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
//
// 	return true
// }

// 方法二：在原字符串上直接判断
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !judge(s[left]) {
			left++
		}
		for left < right && !judge(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func judge(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
