package main

// 签到：矩阵

func maximumWealth(accounts [][]int) int {
	maxRes := 0
	for i := 0; i < len(accounts); i++ {
		demo := 0
		for j := 0; j < len(accounts[i]); j++ {
			demo += accounts[i][j]
		}
		maxRes = maxMaximumWealth(maxRes, demo)
	}
	return maxRes
}

func maxMaximumWealth(a, b int) int {
	if a > b {
		return a
	}
	return b
}
