package main

// 字符串：哈希表，一一对应
// 力扣：https://leetcode.cn/problems/isomorphic-strings/

func isIsomorphic(s string, t string) bool {
	// 要考虑 s是否对应t，t是否对应s
	return judgeST(s, t) && judgeST(t, s)
}

func judgeST(s string, t string) bool {
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

// 或者：
// func isIsomorphic(s, t string) bool {
// 	s2t := map[byte]byte{}
// 	t2s := map[byte]byte{}
// 	for i := range s {
// 		x, y := s[i], t[i]
// 		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
// 			return false
// 		}
// 		s2t[x] = y
// 		t2s[y] = x
// 	}
// 	return true
// }
