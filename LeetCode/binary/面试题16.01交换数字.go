package main

// 异或
// 力扣：https://leetcode.cn/problems/swap-numbers-lcci/

func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] ^ numbers[1]
	numbers[1] = numbers[0] ^ numbers[1] // numbers[1]  = numbers[0] ^ numbers[1] = numbers[0] ^ numbers[1] ^ numbers[1] = numbers[0]
	numbers[0] = numbers[0] ^ numbers[1] // numbers[0] = numbers[0] ^ numbers[1] = numbers[0] ^ numbers[1] ^ numbers[0] = numbers[1]

	return numbers
}
