package main

// 滑动窗口
// 力扣：https://leetcode.cn/problems/maximum-number-of-occurrences-of-a-substring/description/

// 一句话：越短的子串出现的次数越多，所以只需要处理最短字符串minSize的情况，无需理会maxSize
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	mm := make(map[string]int)

	m := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		m[s[right]]++

		if right >= minSize-1 {
			left := right - (minSize - 1)

			if len(m) <= maxLetters {
				mm[s[left:right+1]]++
			}
			m[s[left]]--
			if m[s[left]] == 0 {
				delete(m, s[left])
			}
		}
	}

	maxResult := 0
	for _, v := range mm {
		if maxResult < v {
			maxResult = v
		}
	}
	return maxResult
}

// func main() {
// 	fmt.Println(maxFreq("aababcaab", 2, 3, 4))                                                                                             // 2
// 	fmt.Println(maxFreq("aabcabcab", 2, 2, 3))                                                                                             // 3
// 	fmt.Println(maxFreq("abcde", 2, 3, 3))                                                                                                 // 0
// 	fmt.Println(maxFreq("eddcdfcfbbbdedbcddebbfbbdddacfecddacabdbddabfffdecebcabfbdeadecffbcdccecdebbaaaeebefbabeaefbebecdfcd", 1, 5, 24)) // 0
// }
