package main

// 矩阵旋转
// 力扣：https://leetcode.cn/problems/determine-whether-matrix-can-be-obtained-by-rotation/

func findRotation(mat [][]int, target [][]int) bool {
	for i := 0; i < 4; i++ {
		mat = xuanzhaunFindRotation(mat)
		if judgeFindRotation(mat, target) {
			return true
		}
	}
	return false
}

// 将一个矩阵顺时针选择90度
func xuanzhaunFindRotation(sli [][]int) [][]int {
	n := len(sli)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[j][n-i-1] = sli[i][j]
		}
	}
	return res
}

// 判断两个矩阵是否相同
func judgeFindRotation(sli1, sli2 [][]int) bool {
	n := len(sli1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if sli1[i][j] != sli2[i][j] {
				return false
			}
		}
	}
	return true
}
