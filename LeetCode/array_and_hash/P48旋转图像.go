package main

// 矩阵旋转
// 力扣：https://leetcode.cn/problems/rotate-image/description/

// 方法一：使用辅助数组（额外空间）
// func rotate(matrix [][]int) {
// 	res := xuanzhuanRotate(matrix)
//
// 	copy(matrix, res)
// }
//
// func xuanzhuanRotate(sli [][]int) [][]int {
// 	n := len(sli)
// 	res := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		res[i] = make([]int, n)
// 	}
//
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			res[j][n-i-1] = sli[i][j]
// 		}
// 	}
// 	return res
// }

// 方法二：用翻转代替旋转
// func rotate(matrix [][]int) {
// 	n := len(matrix)
// 	// 根据水平轴翻转
// 	for i := 0; i < n/2; i++ {
// 		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
// 	}
//
// 	// 根据主对角线翻转
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < i; j++ {
// 			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
// 		}
// 	}
// }

// 方法三：原地修改（由外到内一层一层旋转）
// 题解：https://leetcode.cn/problems/rotate-image/solutions/1228078/48-xuan-zhuan-tu-xiang-fu-zhu-ju-zhen-yu-jobi/
func rotate(matrix [][]int) {
	n := len(matrix)

	// 要循环的层数
	for i := 0; i < n/2; i++ {
		// 调整该层的位置
		for j := 0; j < (n+1)/2; j++ {
			temp := matrix[i][j]                    // 存储左上角
			matrix[i][j] = matrix[n-j-1][i]         // 左上角 = 左下角
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1] // 左下角 = 右下角
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1] // 右下角 = 右上角
			matrix[j][n-i-1] = temp                 // 右上角 = 左上角
		}
	}
}
