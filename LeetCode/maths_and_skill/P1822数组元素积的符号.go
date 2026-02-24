package main

// 签到：数学

func arraySign(nums []int) int {
	res := 1
	for _, v := range nums {
		res *= v

		if res > 0 {
			res = 1
		} else if res < 0 {
			res = -1
		} else {
			return 0
		}
	}
	return res
}
