package main

import "bytes"

// 签到：字符串，bytes.Buffer

func mergeAlternately(word1 string, word2 string) string {
	var result bytes.Buffer
	n1, n2 := 0, 0
	for n1 < len(word1) || n2 < len(word2) {
		if n1 < len(word1) {
			result.WriteByte(word1[n1])
			n1++
		}
		if n2 < len(word2) {
			result.WriteByte(word2[n2])
			n2++
		}
	}
	return result.String()
}
