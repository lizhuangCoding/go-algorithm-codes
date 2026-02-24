package main

// 矩阵，二分查找
// 力扣：https://leetcode.cn/problems/search-a-2d-matrix/
// 相似题目（但不同）：https://leetcode.cn/problems/search-a-2d-matrix-ii/description/

// 方法一：使用两次二分查找（找行、找列）（可以用方法三来替代）
// 题解：https://leetcode.cn/problems/search-a-2d-matrix/solutions/688117/sou-suo-er-wei-ju-zhen-by-leetcode-solut-vxui/
// 时间复杂度：O(logm+logn)=O(logmn)，其中 m 和 n 分别是矩阵的行数和列数。
// 空间复杂度：O(1)。
// func searchMatrix(matrix [][]int, target int) bool {
// 	row := sort.Search(len(matrix), func(i int) bool {
// 		return matrix[i][0] > target
// 	}) - 1
//
// 	if row < 0 {
// 		return false
// 	}
// 	col := sort.SearchInts(matrix[row], target)
// 	return col < len(matrix[row]) && matrix[row][col] == target
// }

// 方法二：一次二分查找
// 时间复杂度：O(log(mn))，其中 m 和 n 分别为 matrix 的行数和列数。
// 空间复杂度：O(1)。
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	n, m := len(matrix), len(matrix[0])

	left, right := 0, n*m-1

	for left <= right {
		mid := (right + left) / 2

		x := mid / m
		y := mid % m

		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 方法三：排除法（从右上角缩小到左下角）
// 题解：https://leetcode.cn/problems/search-a-2d-matrix/
// 时间复杂度：O(m+n)，其中 m 和 n 分别为 matrix 的行数和列数。
// 空间复杂度：O(1)。
// func searchMatrix(matrix [][]int, target int) bool {
// 	m, n := len(matrix), len(matrix[0])
// 	i, j := 0, n-1
// 	for i < m && j >= 0 { // 还有剩余元素
// 		if matrix[i][j] == target {
// 			return true // 找到 target
// 		}
// 		if matrix[i][j] < target {
// 			i++ // 这一行剩余元素全部小于 target，排除
// 		} else {
// 			j-- // 这一列剩余元素全部大于 target，排除
// 		}
// 	}
// 	return false
// }
