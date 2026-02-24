package main

// 哈希
// 力扣：https://leetcode.cn/problems/equal-row-and-column-pairs/

// 错误写法：不能用string存储，因为出现这种情况会出问题：[[11,1],[1,11]]
// func equalPairs(grid [][]int) int {
// 	// 存储行
// 	m := make(map[string]int)
// 	for i := 0; i < len(grid); i++ {
// 		demo := ""
// 		for j := 0; j < len(grid[i]); j++ {
// 			demo = fmt.Sprintf("%s%d", demo, grid[i][j])
// 		}
// 		m[demo]++
// 	}
// 	fmt.Println(m)
//
// 	result := 0
// 	// 遍历列，判断是否有对应的行
// 	for i := 0; i < len(grid); i++ {
// 		demo := ""
// 		for j := 0; j < len(grid[0]); j++ {
// 			demo = fmt.Sprintf("%s%d", demo, grid[j][i])
// 		}
//
// 		if k, ok := m[demo]; ok {
// 			fmt.Println(demo)
// 			result += k
// 		}
// 	}
//
// 	return result
// }

func equalPairs(grid [][]int) int {
	m := map[[200]int]int{}
	result := 0

	for _, row := range grid {
		a := [200]int{}
		for j, v := range row {
			a[j] = v
		}
		m[a]++
	}

	for j := range grid[0] {
		a := [200]int{}
		for i, row := range grid {
			a[i] = row[j]
		}

		result += m[a]
	}

	return result
}
