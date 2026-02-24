package main

// 滑动窗口：二、不定长 2.2求最短/最小
// 力扣：https://leetcode.cn/problems/minimum-window-substring/description/

func minWindow(s string, t string) string {
	var cntS, cntT [128]int
	for _, v := range t {
		cntT[v]++
	}

	// 注意范围
	resultLeft, resultRight := -1, len(s)
	left := 0
	for right := 0; right < len(s); right++ {
		cntS[s[right]]++

		for isCoveredMinWindow(cntS, cntT) {
			if resultRight-resultLeft > right-left {
				resultLeft, resultRight = left, right
			}
			cntS[s[left]]--
			left++
		}
	}

	if resultLeft < 0 {
		return ""
	}
	return s[resultLeft : resultRight+1]
}

// 判断切片s是否包含切片t的全部元素
func isCoveredMinWindow(cntS, cntT [128]int) bool {
	for i := 0; i < len(cntS); i++ {
		if cntS[i] < cntT[i] {
			return false
		}
	}
	return true
}
