package main

// 栈
// 力扣：https://leetcode.cn/problems/removing-stars-from-a-string/

func removeStars(s string) string {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
