package main

// 矩阵
// 力扣：https://leetcode.cn/problems/spiral-matrix/description/

func spiralOrder(matrix [][]int) []int {
	rows, columns := len(matrix), len(matrix[0])
	result := make([]int, 0)
	left, right, top, bottom := 0, columns-1, 0, rows-1

	// for left <= right && top <= bottom {
	// 	// 从左到右
	// 	for j := left; j <= right; j++ {
	// 		result = append(result, matrix[top][j])
	// 	}
	//
	// 	// 从上到下
	// 	for i := top + 1; i <= bottom; i++ {
	// 		result = append(result, matrix[i][right])
	// 	}
	//
	// 	if top < bottom && left < right { // 这里不能等于
	// 		// 从右到左
	// 		for j := right - 1; j > left; j-- {
	// 			result = append(result, matrix[bottom][j])
	// 		}
	// 		// 从下到上
	// 		for i := bottom; i > top; i-- {
	// 			result = append(result, matrix[i][left])
	// 		}
	// 	}
	//
	// 	left++
	// 	right--
	// 	top++
	// 	bottom--
	// }

	// 或者
	for left <= right && top <= bottom {
		// 从左到右
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		// 从上到下
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// 从右到左
		if top <= bottom {
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
		}
		bottom--

		// 从下到上
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
	}
	return result
}

// func main() {
// 	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))             // [1,2,3,6,9,8,7,4,5]
// 	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})) // [1,2,3,4,8,12,11,10,9,5,6,7]
// 	fmt.Println(spiralOrder([][]int{{1}}))                                         // 1
// 	fmt.Println(spiralOrder([][]int{{1, 2}}))                                      // 1,2
// 	fmt.Println(spiralOrder([][]int{{1}, {2}}))                                    // 1,2
// }
