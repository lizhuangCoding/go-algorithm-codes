package main

// 字符串：数学，可以看作26进制转换
// 力扣：https://leetcode.cn/problems/excel-sheet-column-title/

// 从后往前算
// func convertToTitle(columnNumber int) string {
// 	s := ""
// 	for columnNumber > 0 {
// 		b := byte(columnNumber % 26)
// 		if b == 0 {
// 			b = 26
// 		}
//
// 		// columnNumber 只剩下26了
// 		if columnNumber == 26 {
// 			s = "Z" + s
// 			return s
// 		} else {
// 			s = string('A'+b-1) + s
// 		}
//
// 		if b == 26 {
// 			columnNumber -= 26 // 记得 columnNumber-26
// 		}
// 		columnNumber /= 26
// 	}
//
// 	return s
// }

// 优化：
func convertToTitle(columnNumber int) string {
	s := ""
	for columnNumber > 0 {
		columnNumber--
		s = string(byte(columnNumber%26)+'A') + s
		columnNumber /= 26
	}
	return s
}
