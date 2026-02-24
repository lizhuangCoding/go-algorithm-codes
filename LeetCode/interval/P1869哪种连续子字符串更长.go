package main

// 分组循环，字符串
// 力扣：https://leetcode.cn/problems/longer-contiguous-segments-of-ones-than-zeros/

func checkZeroOnes(s string) bool {
	n0, n1 := 0, 0
	for i := 0; i < len(s); i++ {
		start := i
		for ; i < len(s)-1 && s[i] == s[i+1]; i++ {
		}
		if s[i] == '0' {
			n0 = checkZeroOnesMax(n0, i-start+1)
		} else {
			n1 = checkZeroOnesMax(n1, i-start+1)
		}
	}
	return n1 > n0
}

func checkZeroOnesMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
