package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/get-equal-substrings-within-budget/description/

func equalSubstring(s string, t string, maxCost int) int {
	maxLen := 0
	left := 0
	demo := 0 // 开销

	for right := 0; right < len(s); right++ {
		d := int(s[right]) - int(t[right])
		if d < 0 {
			d = -d
		}
		// fmt.Println(d)
		demo += d

		for left <= right && demo > maxCost {
			d := int(s[left]) - int(t[left])
			if d < 0 {
				d = -d
			}
			demo -= d

			left++
		}

		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// func main() {
// 	fmt.Println(equalSubstring("abcd", "bcdf", 3))      // 3
// 	fmt.Println(equalSubstring("krrgw", "zjxss", 19))   // 2
// 	fmt.Println(equalSubstring("pxezla", "loewbi", 25)) // 4
// 	fmt.Println(equalSubstring("abcd", "cdef", 1))      // 0
// }
