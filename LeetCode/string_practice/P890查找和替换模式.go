package main

// 字符串：哈希表，一一对应
// 力扣：https://leetcode.cn/problems/find-and-replace-pattern/

func findAndReplacePattern(words []string, pattern string) []string {
	sli := make([]string, 0)
	for i := range words {
		if judgeWordPattern(words[i], pattern) && judgeWordPattern(pattern, words[i]) {
			sli = append(sli, words[i])
		}
	}

	return sli
}

func judgeWordPattern(s string, t string) bool {
	m := make(map[int32]int32)
	for i, v := range s {
		temp, ok := m[v]
		if ok {
			if temp == int32(t[i]) {
				continue
			} else {
				return false
			}
		} else {
			m[v] = int32(t[i])
		}
	}
	return true
}
