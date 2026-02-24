package main

// 动态规划
// 力扣：https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/description/

func trainWays(num int) int {
	dp := make([]int, num+1)
	dp[0] = 1
	if num == 0 {
		return 1
	}
	dp[1] = 1
	if num == 1 {
		return 1
	}
	dp[2] = 2
	if num == 2 {
		return 2
	}

	for i := 3; i <= num; i++ {
		dp[i] = dp[i-2]%1000000007 + dp[i-1]%1000000007
	}
	return dp[num] % 1000000007
}

// 或者：
// func trainWays(num int) int {
// 	a, b := 1, 1
// 	sum := 0
//
// 	for i := 0; i < num; i++ {
// 		sum = (a + b) % 1000000007
//
// 		a = b
// 		b = sum
// 	}
// 	return a
// }
