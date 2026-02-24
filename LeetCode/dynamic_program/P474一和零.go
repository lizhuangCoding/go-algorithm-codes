package main

// 动态规划：01背包、三维
// 力扣：https://leetcode.cn/problems/ones-and-zeroes/description/

// 三维dp
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := 1; i <= len(strs); i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {

				a, b := get0and1(strs[i-1])
				if j < a || k < b {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-a][k-b]+1)
				}

			}
		}
	}

	return dp[len(strs)][m][n]
}

// 判断某个字符串中有多少个0和1
func get0and1(s string) (a int, b int) {
	for _, v := range s {
		if v == '0' {
			a++
		} else {
			b++
		}
	}
	return
}

// 二维dp
// func findMaxForm(strs []string, m int, n int) int {
// 	dp := make([][]int, m+1)
// 	for i := range dp {
// 		dp[i] = make([]int, n+1)
// 	}
//
//
// 	for i := 0; i < len(strs); i++ {
//
// 		zeroNum, oneNum := 0, 0
// 		// 计算0,1 个数
// 		// 或者直接strings.Count(strs[i],"0")
// 		for _, v := range strs[i] {
// 			if v == '0' {
// 				zeroNum++
// 			}
// 		}
// 		oneNum = len(strs[i]) - zeroNum
//
// 		// 从后往前 遍历背包容量
// 		for j := m; j >= zeroNum; j-- {
// 			for k := n; k >= oneNum; k-- {
//
// 				dp[j][k] = max(dp[j][k], dp[j-zeroNum][k-oneNum]+1)
//
// 			}
// 		}
// 	}
//
// 	return dp[m][n]
// }

// func main() {
// 	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)) // 4
// 	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))                   // 2
// }
