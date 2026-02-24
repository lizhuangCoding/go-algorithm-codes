package main

// 签到：字符串
func maxScore(s string) int {
	left0, right1 := 0, 0
	if s[0] == '0' {
		left0 = 1
	}
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			right1++
		}
	}
	maxRes := left0 + right1

	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			left0++
		} else if s[i] == '1' {
			right1--
		}
		maxRes = maxMaxScore(maxRes, left0+right1)
	}
	return maxRes
}

func maxMaxScore(a, b int) int {
	if a > b {
		return a
	}
	return b
}
