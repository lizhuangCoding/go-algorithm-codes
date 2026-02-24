package main

// 哈希，位运算
// 力扣：https://leetcode.cn/problems/repeated-dna-sequences/description/

// 方法一：哈希表map
// 时间复杂度：O(NL)，其中 N 是字符串 s 的长度，L=10 即目标子串的长度。
// 空间复杂度：O(NL)。
// func findRepeatedDnaSequences(s string) []string {
// 	if len(s) <= 10 {
// 		return []string{}
// 	}
//
// 	m := make(map[string]int)
// 	result := make([]string, 0)
//
// 	for i := 0; i <= len(s)-10; i++ {
// 		sub := s[i : i+10]
// 		m[sub]++
// 		if m[sub] == 2 {
// 			result = append(result, sub)
// 		}
// 	}
//
// 	return result
// }

// 方法二：哈希表 + 滑动窗口 + 位运算
// 题解：力扣：https://leetcode.cn/problems/repeated-dna-sequences/solutions/1035568/zhong-fu-de-dnaxu-lie-by-leetcode-soluti-z8zn/
// 时间复杂度：O(N)，其中 N 是字符串 s 的长度。
// 空间复杂度：O(N)。

// 不想看...
