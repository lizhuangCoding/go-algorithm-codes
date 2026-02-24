package main

// 动态规划 一、入门DP 1.1 爬楼梯
// 力扣：https://leetcode.cn/problems/climbing-stairs/description/

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		// 首先是dp[i - 1]，上i-1层楼梯，有dp[i - 1]种方法，那么再一步跳一个台阶不就是dp[i]了么。
		// 还有就是dp[i - 2]，上i-2层楼梯，有dp[i - 2]种方法，那么再一步跳两个台阶不就是dp[i]了么。
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 空间优化：
// func climbStairs(n int) int {
// 	f0, f1 := 1, 1
// 	for i := 2; i <= n; i++ {
// 		f0, f1 = f1, f1+f0
// 	}
// 	return f1
// }

// func main() {
// 	fmt.Println(climbStairs(2)) // 2
// 	fmt.Println(climbStairs(3)) // 3
// 	fmt.Println(climbStairs(4)) // 5
// }
