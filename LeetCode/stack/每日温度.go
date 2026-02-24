package main

// 单调栈（求下一个更大的元素，就是用单调栈解）
// 单调栈的主要特点是，在入栈时保持栈内元素的单调性。具体来说，单调递增栈是指从栈底到栈顶元素递增，而单调递减栈是指从栈底到栈顶元素递减
// 暴力1：超时
func dailyTemperatures(temperatures []int) []int {
	var result = make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {

		j := i
		for ; j < len(temperatures); j++ {
			if temperatures[i] < temperatures[j] {
				break
			}
		}

		if j-i != len(temperatures)-i {
			result[i] = j - i
		} else {
			result[i] = 0
		}

	}
	return result
}
