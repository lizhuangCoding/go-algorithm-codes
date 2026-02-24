package main

import "fmt"

// 滑动窗口
// 力扣：https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description/

// 速度太慢：
func findSubstring(s string, words []string) []int {
	result := make([]int, 0)
	// 每一次滑动窗口的总长度
	l := len(words) * len(words[0])

	fmt.Println(l, len(s)-l)
	for left := 0; left <= len(s)-l; left++ {
		m := make(map[string]int)
		for _, v := range words {
			m[v]++
		}

		// fmt.Println(left, left+l-len(words[0]))

		for i := left; i <= left+l-len(words[0]); i += len(words[0]) {
			if _, ok := m[s[i:i+len(words[0])]]; ok {
				m[s[i:i+len(words[0])]]--
				if m[s[i:i+len(words[0])]] == 0 {
					delete(m, s[i:i+len(words[0])])
				}
			} else {
				break
			}

			if len(m) == 0 {
				result = append(result, left)
			}
		}
	}
	return result
}

// func main() {
// 	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
// 	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
// 	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
// 	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
// }

// func findSubstring(s string, words []string) (ans []int) {
// 	ls, m, n := len(s), len(words), len(words[0])
// 	for i := 0; i < n && i+m*n <= ls; i++ {
// 		differ := map[string]int{}
// 		for j := 0; j < m; j++ {
// 			differ[s[i+j*n:i+(j+1)*n]]++
// 		}
// 		for _, word := range words {
// 			differ[word]--
// 			if differ[word] == 0 {
// 				delete(differ, word)
// 			}
// 		}
//
// 		// 滑动
// 		for start := i; start < ls-m*n+1; start += n {
// 			if start != i {
// 				word := s[start+(m-1)*n : start+m*n]
// 				differ[word]++
// 				if differ[word] == 0 {
// 					delete(differ, word)
// 				}
// 				word = s[start-n : start]
// 				differ[word]--
// 				if differ[word] == 0 {
// 					delete(differ, word)
// 				}
// 			}
// 			if len(differ) == 0 {
// 				ans = append(ans, start)
// 			}
// 		}
// 	}
// 	return
// }
