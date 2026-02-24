package main

// 字符串：自定义构造回文字符串
// 力扣：https://leetcode.cn/problems/longest-palindrome/

func longestPalindrome(s string) int {
	m := make(map[byte]int)
	n := 0

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for _, v := range m {
		n += v / 2 * 2 // 加偶数
	}

	// 加一个中间的字母
	if n < len(s) {
		n++
	}
	return n
}
