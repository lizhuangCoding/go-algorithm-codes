package main

// 技巧：最大公约数
// 力扣：https://leetcode.cn/problems/greatest-common-divisor-of-strings/description/
//
// // 方法一：暴力
// func gcdOfStrings(str1 string, str2 string) string {
// 	len1 := len(str1)
// 	len2 := len(str2)
//
// 	for i := minGcdOfStrings(len1, len2); i >= 1; i-- {
// 		if len1%i == 0 && len2%i == 0 {
// 			X := str1[:i]
// 			// X 必须满足：str1 = X + ... + X 和 str2 = X + ... + X
// 			if checkGcdOfStrings(X, str1) && checkGcdOfStrings(X, str2) {
// 				return X
// 			}
// 		}
// 	}
// 	return ""
// }
//
// func checkGcdOfStrings(t string, s string) bool {
// 	lenT := len(t)
// 	lenS := len(s)
// 	if lenS%lenT != 0 {
// 		return false
// 	}
//
// 	repeatCount := lenS / lenT
// 	ans := ""
// 	for i := 0; i < repeatCount; i++ {
// 		ans += t
// 	}
// 	return ans == s
// }
//
// func minGcdOfStrings(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// 方法二：最大公约数
func gcdOfStrings(str1 string, str2 string) string {
	// 假设str1是N个x，str2是M个x，那么str1+str2肯定是等于str2+str1的。
	if str1+str2 != str2+str1 {
		return ""
	}

	// 找到最大公约数
	index := gcdGcdOfStrings(len(str1), len(str2))
	return str1[:index]
}

// 辗转相除法求gcd
func gcdGcdOfStrings(a, b int) int {
	for b != 0 {
		demo := b
		b = a % b
		a = demo
	}
	return a
}
