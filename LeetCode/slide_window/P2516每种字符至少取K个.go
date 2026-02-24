package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/take-k-of-each-character-from-left-and-right/

// 逆向思维：找到可以保留的最大长度
func takeCharacters(s string, k int) int {
	a, b, c := -k, -k, -k // 需要保留的abc个数
	for _, v := range s {
		if v == 'a' {
			a++
		} else if v == 'b' {
			b++
		} else {
			c++
		}
	}
	// 字母不够
	if a < 0 || b < 0 || c < 0 {
		return -1
	}

	maxLen := 0
	left := 0
	numA, numB, numC := 0, 0, 0
	// 找到最大长度的字符串：numA<=a，numB<=b，numC<=c
	for right := 0; right < len(s); right++ {
		if s[right] == 'a' {
			numA++
		} else if s[right] == 'b' {
			numB++
		} else {
			numC++
		}
		// 如果超出a，b，c一定错误
		for numA > a || numB > b || numC > c {
			if s[left] == 'a' {
				numA--
			} else if s[left] == 'b' {
				numB--
			} else {
				numC--
			}
			left++
		}
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}
	return len(s) - maxLen
}

// 优化：
// func takeCharacters(s string, k int) int {
// 	cnt := [3]int{}
// 	for _, c := range s {
// 		cnt[c-'a']++ // 一开始，把所有字母都取走
// 	}
// 	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
// 		return -1 // 字母个数不足 k
// 	}
//
// 	mx, left := 0, 0
// 	for right, c := range s {
// 		c -= 'a'
// 		cnt[c]--         // 移入窗口，相当于不取走 c
// 		for cnt[c] < k { // 窗口之外的 c 不足 k
// 			cnt[s[left]-'a']++ // 移出窗口，相当于取走 s[left]
// 			left++
// 		}
// 		if mx < right-left+1 {
// 			mx = right - left + 1
// 		}
// 	}
// 	return len(s) - mx
// }
