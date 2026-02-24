package main

// 字符串
// 力扣：https://leetcode.cn/problems/first-unique-character-in-a-string/description/

func firstUniqChar(s string) int {
	var a = [26]int{}

	for i := 0; i < len(s); i++ {
		a[s[i]-'a']++
	}

	for i := range s {
		if a[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// 或者使用队列
