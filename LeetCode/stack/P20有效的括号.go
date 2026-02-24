package main

// 栈
// 力扣：https://leetcode.cn/problems/valid-parentheses/description/

func isValid(s string) bool {
	// 长度必须是偶数
	if len(s)%2 != 0 {
		return false
	}

	sli := make([]byte, 0)

	m := make(map[byte]byte)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['

	for i := 0; i < len(s); i++ {
		demo := s[i]

		d, ok := m[demo]
		if ok { // 如果是 ) ] }，判断栈中是否有对应的值，如果没有对应的值，那么就返回false

			if len(sli) >= 1 && sli[len(sli)-1] == d {
				sli = sli[:len(sli)-1]
			} else {
				return false
			}

		} else { // 继续添加到栈中
			sli = append(sli, demo)
		}

	}

	return len(sli) == 0
}
