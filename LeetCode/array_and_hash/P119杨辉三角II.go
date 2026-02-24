package main

// 数组
// 力扣：https://leetcode.cn/problems/pascals-triangle-ii/description/

// 方法一：暴力（求出整个杨辉三角形）
// func getRow(rowIndex int) []int {
// 	result := make([][]int, rowIndex+1)
// 	for i := 0; i < rowIndex+1; i++ {
// 		result[i] = make([]int, i+1)
// 	}
//
// 	for i := 0; i < len(result); i++ {
// 		for j := 0; j < len(result[i]); j++ {
// 			if j == 0 || j == i {
// 				result[i][j] = 1
// 			} else {
// 				result[i][j] = result[i-1][j-1] + result[i-1][j]
// 			}
// 		}
// 	}
// 	return result[rowIndex]
// }

// 优化：注意到对第 i+1 行的计算仅用到了第 i 行的数据，因此可以使用滚动数组的思想优化空间复杂度。
// func getRow(rowIndex int) []int {
// 	var pre, cur []int
// 	for i := 0; i < rowIndex+1; i++ {
// 		cur = make([]int, i+1)
// 		cur[0] = 1
// 		cur[i] = 1
// 		for j := 1; j < i; j++ {
// 			cur[j] = pre[j-1] + pre[j]
// 		}
// 		pre = cur
// 	}
// 	return cur
// }

// 还没看懂
// 优化：当前行第 i 项的计算只与上一行第 i−1 项及第 i 项有关。因此我们可以倒着计算当前行，这样计算到第 i 项时，第 i−1 项仍然是上一行的值。
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}

// 方法二：线性递推
// func getRow(rowIndex int) []int {
// 	row := make([]int, rowIndex+1)
// 	row[0] = 1
// 	for i := 1; i <= rowIndex; i++ {
// 		row[i] = row[i-1] * (rowIndex - i + 1) / i
// 	}
// 	return row
// }

// 方法三：
