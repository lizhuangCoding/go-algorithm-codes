package main

// 字符串：数学，可以看作26进制转换
// 力扣：https://leetcode.cn/problems/excel-sheet-column-number/description/

// ZY = ('Z'-'A'+1) * 26^1 + ('Y'-'A'+1) * 26^0 = 26 * 26 + 25 * 1 = 701
// AB = 1 * 26^1 + 2 * 26^0
// func titleToNumber(columnTitle string) int {
// 	n := 0
// 	l := len(columnTitle) - 1
// 	for i, v := range columnTitle {
// 		n += PowTitleToNumber(26, l-i) * (int(v - 'A' + 1))
// 	}
//
// 	return n
// }
// // PowTitleToNumber 计算a的b次方
// func PowTitleToNumber(a, b int) int {
// 	result := 1
// 	for i := 0; i < b; i++ {
// 		result *= a
// 	}
// 	return result
// }

// 优化：
func titleToNumber(columnTitle string) int {
	res := 0
	for i := range columnTitle {
		res = res*26 + int(columnTitle[i]-'A') + 1
	}
	return res
}
