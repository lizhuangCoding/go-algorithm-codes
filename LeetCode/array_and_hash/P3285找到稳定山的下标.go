package main

// 签到：数组

func stableMountains(height []int, threshold int) []int {
	result := make([]int, 0)
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			result = append(result, i)
		}
	}
	return result
}
