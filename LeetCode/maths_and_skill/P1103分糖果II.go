package main

// 数学

// 方法一：模拟
func distributeCandies(candies, n int) []int {
	ans := make([]int, n)
	for i := 1; candies > 0; i++ {
		ans[(i-1)%n] += minDistributeCandies(i, candies)
		candies -= i
	}
	return ans
}

func minDistributeCandies(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法二：数学...
