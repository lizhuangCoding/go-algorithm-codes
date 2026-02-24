package main

import "math"

// 滑动窗口：二、不定长 2.2求最短/最小
// 力扣：https://leetcode.cn/problems/replace-the-substring-for-balanced-string/description/

func balancedString(s string) int {
	n := len(s)
	target := n / 4
	Q, W, E, R := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'Q' {
			Q++
		} else if s[i] == 'W' {
			W++
		} else if s[i] == 'E' {
			E++
		} else if s[i] == 'R' {
			R++
		}
	}

	// 已经符合要求啦
	if Q == target && W == target && E == target && R == target {
		return 0
	}

	result := math.MaxInt64
	left := 0
	q, w, e, r := 0, 0, 0, 0

	// 核心：如果在待替换子串之外的任意字符的出现次数超过 n/4，那么无论怎么替换，都无法使这个字符在整个字符串中的出现次数为 n/4。
	for right := 0; right < n; right++ {
		if s[right] == 'Q' {
			q++
		} else if s[right] == 'W' {
			w++
		} else if s[right] == 'E' {
			e++
		} else if s[right] == 'R' {
			r++
		}

		// 如果在待替换子串之外的任意字符的出现次数都 <= n/4，那么可以通过替换，使 s 为平衡字符串，即每个字符的出现次数均为 n/4。
		for Q-q <= target && W-w <= target && E-e <= target && R-r <= target {
			if result > right-left+1 {
				result = right - left + 1
			}

			if s[left] == 'Q' {
				q--
			} else if s[left] == 'W' {
				w--
			} else if s[left] == 'E' {
				e--
			} else if s[left] == 'R' {
				r--
			}
			left++ // 缩小子串
		}
	}
	return result
}

// 简洁度优化：
// func balancedString(s string) int {
// 	cnt, m := ['X']int{}, len(s)/4 // 也可以用哈希表，不过数组更快一些
// 	for _, c := range s {
// 		cnt[c]++
// 	}
// 	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m {
// 		return 0 // 已经符合要求啦
// 	}
// 	ans, left := len(s), 0
// 	for right, c := range s { // 枚举子串右端点
// 		cnt[c]--
// 		for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m {
// 			ans = min(ans, right-left+1)
// 			cnt[s[left]]++
// 			left++ // 缩小子串
// 		}
// 	}
// 	return ans
// }
