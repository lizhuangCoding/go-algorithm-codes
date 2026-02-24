package main

// 模拟
// 力扣：https://leetcode.cn/problems/robot-bounded-in-circle/description/
// 题解：https://leetcode.cn/problems/robot-bounded-in-circle/

// 机器人想要摆脱循环，在一串指令之后的状态，必须是不位于原点且方向朝北。
func isRobotBounded(instructions string) bool {
	direc := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 分别是朝向 北、东、南、西
	direcIndex := 0                                    // 初始为0，朝向北
	x, y := 0, 0

	// 根据不同类型的指令调整机器人的位置或方向
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		if instruction == 'G' { // 调整位置，向当前方向移动一步
			x += direc[direcIndex][0]
			y += direc[direcIndex][1]
		} else if instruction == 'L' { // 调整方向
			// direcIndex = (direcIndex - 1 + 4) % 4 等价于 direcIndex += 3; direcIndex %= 4。使用 +3 避免直接减法出现负索引。
			direcIndex += 3
			direcIndex %= 4
		} else { // 调整方向
			direcIndex++
			direcIndex %= 4
		}
	}
	// direcIndex != 0:如果机器人执行完所有指令后，方向不是初始方向（朝北），则重复相同指令会形成循环轨迹。
	// (x == 0 && y == 0):如果机器人回到了原点 (0, 0)，也会形成循环轨迹。
	return direcIndex != 0 || (x == 0 && y == 0)
}
