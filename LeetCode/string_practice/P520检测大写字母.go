package main

import "unicode"

// 字符串
// 力扣：https://leetcode.cn/problems/detect-capital/description/

// func detectCapitalUse(word string) bool {
//
// 	if judgeDetectCapitalUse1(word) || judgeDetectCapitalUse2(word) || judgeDetectCapitalUse3(word) {
// 		return true
// 	}
//
// 	return false
// }
//
// // 判断全部大写
// func judgeDetectCapitalUse1(s string) bool {
// 	for _, v := range s {
// 		if v < 'A' || v > 'Z' {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// // 判断全部小写
// func judgeDetectCapitalUse2(s string) bool {
// 	for _, v := range s {
// 		if v < 'a' || v > 'z' {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// // 判断第一个大写其余小写
// func judgeDetectCapitalUse3(s string) bool {
// 	for k, v := range s {
// 		if k == 0 && (v < 'A' || v > 'Z') {
// 			return false
// 		}
//
// 		if k >= 1 && (v < 'a' || v > 'z') {
// 			return false
// 		}
// 	}
// 	return true
// }

// 优化
func detectCapitalUse(word string) bool {
	// 若第 1 个字母为小写，则需额外判断第 2 个字母是否为小写
	if len(word) >= 2 && unicode.IsLower(rune(word[0])) && unicode.IsUpper(rune(word[1])) {
		return false
	}

	// 无论第 1 个字母是否大写，其他字母必须与第 2 个字母的大小写相同
	for i := 2; i < len(word); i++ {
		if unicode.IsLower(rune(word[i])) != unicode.IsLower(rune(word[1])) {
			return false
		}
	}
	return true
}
