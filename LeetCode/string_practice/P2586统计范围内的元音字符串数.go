package main

// 签到：字符串

func vowelStrings(words []string, left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		end := len(words[i]) - 1
		if (words[i][0] == 'a' || words[i][0] == 'e' || words[i][0] == 'i' || words[i][0] == 'o' || words[i][0] == 'u') &&
			(words[i][end] == 'a' || words[i][end] == 'e' || words[i][end] == 'i' || words[i][end] == 'o' || words[i][end] == 'u') {
			res++
		}
	}

	// 或者：
	// for _, s := range words[left : right+1] {
	// 	if strings.Contains("aeiou", s[:1]) && strings.Contains("aeiou", s[len(s)-1:]) {
	// 		res++
	// 	}
	// }
	return res
}
