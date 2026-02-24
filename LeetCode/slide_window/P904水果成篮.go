package main

// 滑动窗口：二、不定长 2.1求最长/最大
// 力扣：https://leetcode.cn/problems/fruit-into-baskets/description/

func totalFruit(fruits []int) int {
	m := make(map[int]int)

	maxNum, left := 0, 0
	for right := 0; right < len(fruits); right++ {
		m[fruits[right]]++

		for left < right && len(m) > 2 {
			m[fruits[left]]--
			if m[fruits[left]] == 0 {
				delete(m, fruits[left])
			}
			left++
		}

		if maxNum < right-left+1 {
			maxNum = right - left + 1
		}
	}
	return maxNum
}
