package main

// 数组
// 力扣：https://leetcode.cn/problems/transpose-matrix/description/

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}

// func main() {
// 	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // [[1,4,7],[2,5,8],[3,6,9]]
// 	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))            // [1,4],[2,5],[3,6]]
// }
