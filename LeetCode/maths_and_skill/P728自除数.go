package main

// 签到：数学

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)

	for i := left; i <= right; i++ {
		demo := i
		b := true
		for demo > 0 {
			if demo%10 == 0 {
				b = false
				break
			}
			if i%(demo%10) != 0 {
				b = false
				break
			}
			demo /= 10
		}
		if b {
			result = append(result, i)
		}
	}
	return result
}
