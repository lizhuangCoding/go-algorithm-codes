package main

// 签到：哈希表

func isPossibleToRearrange(s string, t string, k int) bool {
	m := make(map[string]int)
	l := len(s) / k // 单个子字符串的长度

	for i := 0; i < len(s); i += l {
		m[s[i:i+l]]++
	}
	// fmt.Println(m)

	for i := 0; i < len(t); i += l {
		m[t[i:i+l]]--
		v := m[t[i:i+l]]
		if v == 0 {
			delete(m, t[i:i+l])
		}
	}
	// fmt.Println(m)
	return len(m) == 0
}

// func main() {
// 	fmt.Println(isPossibleToRearrange("abcd", "cdab", 2))     // true
// 	fmt.Println(isPossibleToRearrange("aabbcc", "bbaacc", 3)) // true
// 	fmt.Println(isPossibleToRearrange("aabbcc", "bbaacc", 2)) // false
// }
