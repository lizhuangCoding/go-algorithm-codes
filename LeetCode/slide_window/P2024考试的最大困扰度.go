package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/longest-substring-without-repeating-characters/

func maxConsecutiveAnswers(answerKey string, k int) int {
	t, f := 0, 0
	maxResult := 0
	left := 0

	for right := 0; right < len(answerKey); right++ {
		if answerKey[right] == 'T' {
			t++
		} else {
			f++
		}

		// 如果t和f的个数同时超出了k个，那么就需要--
		for left < right && t > k && f > k {
			if answerKey[left] == 'T' {
				t--
			} else {
				f--
			}
			left++
		}

		if maxResult < right-left+1 {
			maxResult = right - left + 1
		}
	}
	return maxResult
}

// func main() {
// 	fmt.Println(maxConsecutiveAnswers("TTFF", 2))     // 4
// 	fmt.Println(maxConsecutiveAnswers("TFFT", 1))     // 3
// 	fmt.Println(maxConsecutiveAnswers("TTFTTFTT", 1)) // 5
// }
