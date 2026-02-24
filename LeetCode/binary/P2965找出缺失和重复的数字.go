package main

// 位运算
// 力扣：https://leetcode.cn/problems/find-missing-and-repeated-values/

func findMissingAndRepeatedValues(grid [][]int) []int {
	total := 0 // total 是两个结果值的异或（a^b）
	n := len(grid)
	for i := 1; i <= n*n; i++ {
		total ^= i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			total ^= grid[i][j]
		}
	}

	// 找到二进制的最后一位1，根据这个1来区分两个值
	total = total & (-total)
	a, b := 0, 0

	// 找出a，b两个值
	for i := 1; i <= n*n; i++ {
		if total&i == 0 {
			a ^= i
		} else {
			b ^= i
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if total&grid[i][j] == 0 {
				a ^= grid[i][j]
			} else {
				b ^= grid[i][j]
			}
		}
	}

	// 区分a，b
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == a {
				return []int{a, b}
			}
			if grid[i][j] == b {
				return []int{b, a}
			}
		}
	}
	return nil
}
