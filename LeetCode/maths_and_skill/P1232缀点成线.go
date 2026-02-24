package main

// 数学
// 力扣：https://leetcode.cn/problems/check-if-it-is-a-straight-line/description/

func checkStraightLine(coordinates [][]int) bool {
	// 整体移动，让这条线经过原点0
	x, y := coordinates[0][0], coordinates[0][1]
	for _, p := range coordinates {
		p[0] -= x
		p[1] -= y
	}

	// 判断其他点与 点1和点2组成的一条线 是否经过原点
	a, b := coordinates[1][0], coordinates[1][1]
	// Y1 = kX1，Y2 = kX2，所以 Y1*X2 - Y2*X1 = 0
	for _, p := range coordinates {
		if p[0]*b-p[1]*a != 0 {
			return false
		}
	}
	return true
}
