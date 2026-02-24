package main

// 哈希
// 力扣：https://leetcode.cn/problems/determine-if-two-strings-are-close/description/

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	demo1, demo2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(word1); i++ {
		demo1[word1[i]]++
		demo2[word2[i]]++
	}

	// 判断字符种类是否一致
	for k := range demo1 {
		if _, ok := demo2[k]; !ok {
			return false
		}
	}

	m1 := make(map[int]int)
	for _, v := range demo1 {
		m1[v]++
	}

	// 判断字符次数是否一致
	for _, v := range demo2 {
		if _, ok := m1[v]; ok && m1[v] > 0 {
			m1[v]--
		} else {
			return false
		}
	}

	return true
}
