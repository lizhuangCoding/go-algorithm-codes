package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/get-equal-substrings-within-budget/description/

// func longestSemiRepetitiveSubstring(s string) int {
// 	maxLen := 0
// 	left := 1
// 	num := 0 // 拥有几对
// 	for right := 1; right < len(s); right++ {
// 		if s[right-1] == s[right] {
// 			num++
// 		}
//
// 		for left < right && num >= 2 {
// 			if s[left-1] == s[left] {
// 				num--
// 			}
// 			left++
// 		}
//
// 		if maxLen < right-left+1 {
// 			maxLen = right - left + 1
// 		}
// 	}
// 	return maxLen + 1
// }

// 或者：
func longestSemiRepetitiveSubstring(s string) int {
	ans, left, same := 1, 0, 0
	for right := 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			same++

			if same >= 2 {
				left++
				for s[left] != s[left-1] {
					left++
				}
				same--
			}
		}
		if ans < right-left+1 {
			ans = right - left + 1
		}
	}
	return ans
}

// func main() {
// 	fmt.Println(longestSemiRepetitiveSubstring("52233"))
// 	fmt.Println(longestSemiRepetitiveSubstring("5494"))
// 	fmt.Println(longestSemiRepetitiveSubstring("1111111"))
// }
