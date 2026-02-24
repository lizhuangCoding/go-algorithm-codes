package main

// 字符串：模拟
// 力扣：https://leetcode.cn/problems/roman-to-integer/

func romanToInt(s string) int {
	m := make(map[byte]int)
	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000
	n := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[s[i]] < m[s[i+1]] {
			n += m[s[i+1]] - m[s[i]]
			i += 1
		} else {
			n += m[s[i]]
		}
	}

	return n
}

// 或者：
// var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
//
// func romanToInt(s string) (ans int) {
// 	n := len(s)
// 	for i := range s {
// 		value := symbolValues[s[i]]
// 		if i < n-1 && value < symbolValues[s[i+1]] {
// 			ans -= value
// 		} else {
// 			ans += value
// 		}
// 	}
// 	return
// }
