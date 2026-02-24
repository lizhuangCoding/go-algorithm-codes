package main

// 矩阵查找
// 力扣：https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
// 相似题目（但不同）：https://leetcode.cn/problems/search-a-2d-matrix/

// 方法一：二分查找。对每一行都进行二分查找
// 时间复杂度：O(mlogn)。对一行使用二分查找的时间复杂度为 O(logn)，最多需要进行 m 次二分查找。
// 空间复杂度：O(1)。
// func searchMatrix(matrix [][]int, target int) bool {
// 	for _, row := range matrix {
// 		i := sort.SearchInts(row, target)
// 		if i < len(row) && row[i] == target {
// 			return true
// 		}
// 	}
// 	return false
// }

// 方法二：排除法（从右上角不断压缩范围直到左下角）
// 时间复杂度：O(m+n)，其中 m 和 n 分别为 matrix 的行数和列数。
// 空间复杂度：O(1)。
// func searchMatrix(matrix [][]int, target int) bool {
// 	m, n := len(matrix), len(matrix[0])
//
// 	i, j := 0, n-1
//
// 	for i < m && j >= 0 {
// 		if matrix[i][j] == target {
// 			return true
// 		} else if matrix[i][j] < target {
// 			i++
// 		} else {
// 			j--
// 		}
// 	}
// 	return false
// }
