package main

import "strings"

// 签到：数组

func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0)
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], string(x)) {
			res = append(res, i)
		}
	}

	return res
}
