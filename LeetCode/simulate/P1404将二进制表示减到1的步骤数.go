package main

// 模拟
// 力扣：https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/description/

// func numSteps(s string) int {
// 	ans := 0
// 	jinwei := 0
//
// 	for len(s) > 1 {
// 		demo := s[len(s)-1]
// 		if demo == '1' && jinwei == 1 {
// 			// 偶数除以2
// 			jinwei = 1
// 			ans++
// 		} else if demo == '1' && jinwei == 0 {
// 			// 奇数+1，偶数除以2
// 			jinwei = 1
// 			ans += 2
// 		} else if demo == '0' && jinwei == 1 {
// 			// 奇数+1，偶数除以2
// 			jinwei = 1
// 			ans += 2
// 		} else if demo == '0' && jinwei == 0 {
// 			// 偶数除以2
// 			jinwei = 0
// 			ans++
// 		}
// 		s = s[:len(s)-1]
// 	}
//
// 	if jinwei == 1 {
// 		ans++
// 	}
//
// 	return ans
// }

// 优化：
func numSteps(s string) int {
	ans := 0
	jinwei := 0

	for i := len(s) - 1; i > 0; i-- {
		// 将byte转换为int进行运算
		bit := int(s[i]-'0') + jinwei

		if bit == 1 {
			// 奇数：+1（产生进位）然后除以2
			jinwei = 1
			ans += 2
		} else {
			// 偶数：直接除以2（bit为0或2）
			ans++
			// bit为2时，carry保持为1
		}
	}

	// 处理最高位
	if jinwei == 1 {
		ans++
	}

	return ans
}
