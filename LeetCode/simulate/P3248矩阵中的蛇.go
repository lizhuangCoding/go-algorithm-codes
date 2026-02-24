package main

// 签到：模拟

// func finalPositionOfSnake(n int, commands []string) int {
// 	sli := make([][]int, n)
// 	for i := 0; i < n; i++ {
// 		sli[i] = make([]int, n)
// 		for j := 0; j < n; j++ {
// 			sli[i][j] = (i * n) + j
// 		}
// 	}
//
// 	x, y := 0, 0
// 	for _, v := range commands {
// 		if v == "UP" {
// 			x--
// 		} else if v == "RIGHT" {
// 			y++
// 		} else if v == "LEFT" {
// 			y--
// 		} else {
// 			x++
// 		}
// 	}
// 	return sli[x][y]
// }

// 或者：
func finalPositionOfSnake(n int, commands []string) int {
	i, j := 0, 0
	for _, s := range commands {
		switch s[0] {
		case 'U':
			i--
		case 'D':
			i++
		case 'L':
			j--
		default:
			j++
		}
	}
	return i*n + j
}
