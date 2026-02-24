package main

// 栈
// 力扣：https://leetcode.cn/problems/asteroid-collision/description/

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for i := 0; i < len(asteroids); i++ {
		demo := asteroids[i]

		for len(stack) > 0 && demo < 0 && stack[len(stack)-1] > 0 {
			// 冲撞
			if demo+stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
				demo = 0

			} else if demo+stack[len(stack)-1] > 0 {
				demo = 0

			} else {
				stack = stack[:len(stack)-1]
			}
		}

		if demo != 0 {
			stack = append(stack, demo)
		}
	}
	return stack
}
