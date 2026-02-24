package main

// 签到:哈希表

func maximumNumberOfStringPairs(words []string) int {
	m := make(map[string]struct{})
	for _, v := range words {
		m[v] = struct{}{}
	}

	res := 0
	for k, _ := range m {
		delete(m, k)
		demo := reverseMaximumNumberOfStringPairs(k)
		if _, ok := m[demo]; ok {
			res++
		}
	}
	return res
}

// 反转字符串
func reverseMaximumNumberOfStringPairs(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res = res + string(s[i])
	}
	return res
}
