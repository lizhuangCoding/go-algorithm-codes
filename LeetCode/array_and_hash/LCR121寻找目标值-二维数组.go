package main

// 矩阵：从左到右，从上到下依次递增
// 力扣：https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

// 需要从右上角或者左下角开始遍历，然后不断缩短长和宽（不能从左上角或者右下角开始遍历，因为这样不能缩短长和宽）
func findTargetIn2DPlants(plants [][]int, target int) bool {
	if len(plants) == 0 || len(plants[0]) == 0 {
		return false
	}

	m, n := len(plants), len(plants[0]) // 长和宽

	// 剪枝
	if target < plants[0][0] || target > plants[m-1][n-1] {
		return false
	}

	x, y := 0, n-1 // 当前的位置

	// 如果到达 最后一行第一列 那么就再也不可能找到数据了，因为已经遍历结束了
	for x < m && y >= 0 {
		if plants[x][y] == target {
			return true
		} else if plants[x][y] > target { // 左移，消除当前列（可以画个图看一看）
			y--
		} else if plants[x][y] < target { // 下移，消除当前行
			x++
		}
	}

	return false
}

// func main() {
// 	fmt.Println(findTargetIn2DPlants([][]int{{2, 3, 6, 8}, {4, 5, 8, 9}, {5, 9, 10, 12}}, 8)) // true
// 	fmt.Println(findTargetIn2DPlants([][]int{{1, 3, 5}, {2, 5, 7}}, 4))                       // f
// 	fmt.Println(findTargetIn2DPlants([][]int{}, 1))                                           // f
// 	fmt.Println(findTargetIn2DPlants([][]int{{-5}}, -5))                                      // f
// }
