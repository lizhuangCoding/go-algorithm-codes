package main

// 哈希表，计数
// 力扣：https://leetcode.cn/problems/find-the-number-of-winning-players/?envType=daily-question&envId=2024-11-23

func winningPlayerCount(n int, pick [][]int) int {
	m := make(map[int][11]int)
	for _, p := range pick {
		// 从 map 中获取当前键的值
		demo := m[p[0]]
		// 增加对应的位置
		demo[p[1]]++
		// 将更新后的数组存回 map
		m[p[0]] = demo
	}
	// fmt.Println(m)

	ans := 0
	for k, v := range m {
		b := false
		for i := 0; i < len(v); i++ {
			if v[i] >= k+1 { // 只要有一个颜色的数量合格就可以了
				b = true
				break
			}
		}

		if b {
			ans++
		}

	}
	return ans
}

// 或者：二维切片
// func winningPlayerCount(n int, pick [][]int) int {
// 	cnt := make([][]int, n)
// 	for i := range cnt {
// 		cnt[i] = make([]int, 11)
// 	}
// 	for _, p := range pick {
// 		cnt[p[0]][p[1]]++
// 	}
//
// 	ans := 0
// 	for i := 0; i < n; i++ {
// 		for j := 0; j <= 10; j++ {
// 			if cnt[i][j] > i {
// 				ans++
// 				break
// 			}
// 		}
// 	}
// 	return ans
// }

// func main() {
// 	fmt.Println(winningPlayerCount(5, [][]int{{0, 0}, {1, 0}, {1, 0}, {2, 1}, {2, 1}, {2, 0}})) // 2
// 	fmt.Println(winningPlayerCount(5, [][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}}))                 // 0
// 	fmt.Println(winningPlayerCount(5, [][]int{{1, 3}, {1, 3}, {0, 10}, {1, 6}}))                // 2
// }
