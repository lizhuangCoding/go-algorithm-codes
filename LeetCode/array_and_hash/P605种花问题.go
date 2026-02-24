package main

// 数组
// 力扣：https://leetcode.cn/problems/can-place-flowers/description/

// 在最左侧：当前位置为0，下一个位置为0
// 在中间：当前位置为0，上一个位置，下一个位置不为1
// 在最右侧：当前位置为0，上一个位置为0
func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 {
		if (flowerbed[0] == 0 && n == 1) || n == 0 {
			return true
		}
		return false
	}

	res := 0

	for i := 0; i < len(flowerbed); i++ {
		if res == n {
			return true
		}

		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				res++
			}
		} else if i == len(flowerbed)-1 {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 {
				flowerbed[i] = 1
				res++
			}
		} else {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				res++
			}
		}

	}
	return res >= n
}

// 代码优化：
// func canPlaceFlowers(flowerbed []int, n int) bool {
// 	m := len(flowerbed)
// 	for i := 0; i < m; i++ {
// 		if (i == 0 || flowerbed[i-1] == 0) && flowerbed[i] == 0 && (i == m-1 || flowerbed[i+1] == 0) {
// 			n--
// 			i++ // 下一个位置肯定不能种花，直接跳过
// 		}
// 	}
// 	return n <= 0
// }
