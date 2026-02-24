package main

// 签到：数学

func differenceOfSum(nums []int) int {
	num1, num2 := 0, 0
	for i := 0; i < len(nums); i++ {
		demo := nums[i]
		num1 += demo

		for demo > 0 {
			num2 += demo % 10
			demo /= 10
		}
	}

	if num1-num2 < 0 {
		return num2 - num1
	}
	return num1 - num2
}
