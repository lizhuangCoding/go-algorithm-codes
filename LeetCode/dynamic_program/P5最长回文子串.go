package main

// 动态规划
// 力扣：https://leetcode.cn/problems/longest-palindromic-substring/description/
// 题解：https://leetcode.cn/problems/longest-palindromic-substring/solutions/255195/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/?envType=problem-list-v2&envId=7fKmHi2m

func longestPalindrome(s string) string {
	L := len(s)
	if L < 2 {
		return s
	}

	//  dp[i][j] 表示 子串[i,j] 是否是回文串
	dp := make([][]bool, L)
	maxLen := 1
	beginIndex := 0

	// 初始化，单个字符一定是回文字符串
	for i := 0; i < L; i++ {
		dp[i] = make([]bool, L)
		dp[i][i] = true
	}

	// 最外层循环是要查询的回文子串长度
	for l := 2; l <= L; l++ {
		for left := 0; left < L; left++ { // 子串的左边界
			right := left + l - 1 // 子串的右边界

			// 右边界越界，退出当前循环
			if right >= L {
				break
			}

			// 可以直接判定 子串[left,right] 不是回文串
			if s[left] != s[right] {
				dp[left][right] = false
			} else {
				if right-left <= 2 { // 可以直接判定该子串是回文串（加这个判定是为了防止：cbbd 中的 子串 bb 为false）
					dp[left][right] = true
				} else { // 子串[left,right]：left和right位置的元素相同，但是还要判断里面是不是回文串，所以 dp[left][right] 是不是回文串取决于 dp[left+1][right-1]
					dp[left][right] = dp[left+1][right-1]
				}
			}

			// 如果是回文串，尝试更新数据
			if dp[left][right] {
				if right-left+1 > maxLen {
					maxLen = right - left + 1
					beginIndex = left
				}
			}
		}

	}
	return s[beginIndex : beginIndex+maxLen]
}
