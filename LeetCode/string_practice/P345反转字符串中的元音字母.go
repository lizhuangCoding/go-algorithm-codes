package main

// 字符串，双指针
// 力扣：https://leetcode.cn/problems/reverse-vowels-of-a-string/description/

func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	demo := []byte(s)

	for left < right {
		if judgeReverseVowels(demo[left]) && judgeReverseVowels(demo[right]) {
			demo[left], demo[right] = demo[right], demo[left]
			left++
			right--
		}

		if !judgeReverseVowels(demo[left]) {
			left++
		}

		if !judgeReverseVowels(demo[right]) {
			right--
		}
	}

	return string(demo)
}

// 判断是否是元音字母
func judgeReverseVowels(b byte) bool {
	if b == 'a' || b == 'A' || b == 'e' || b == 'E' || b == 'i' || b == 'I' || b == 'o' || b == 'O' || b == 'u' || b == 'U' {
		return true
	}
	return false
}
